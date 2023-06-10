package postgres

import (
	"errors"

	"authentication/src/core/domain/account"
	"authentication/src/core/domain/authorization"
	"authentication/src/core/interfaces/repository"
	"authentication/src/core/messages"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var _ repository.AccountLoader = (*AccountRepository)(nil)

type AccountRepository struct {
	ConnectorManager
}

func (repo AccountRepository) SignUp(userAccount account.Account) (string, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	transaction, err := conn.Beginx()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.TransactionErr)
	}
	defer func() { _ = transaction.Rollback() }()

	registerUser := `insert into account(email, username, full_name, password, hash) 
		values ($1, $2, $3, $4, $5) returning id;`

	var accountID uuid.UUID
	err = transaction.Get(
		&accountID,
		registerUser,
		userAccount.Email(),
		userAccount.Username(),
		userAccount.FullName(),
		userAccount.Password(),
		userAccount.Hash(),
	)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.InsertionErr)
	}

	registerRole := `insert into account_role(account_id,role_id) 
		values ($1, $2);`

	for _, role := range userAccount.Roles() {
		_, err = transaction.Exec(registerRole, accountID, role.ID())
		if err != nil {
			log.Err(err).Msg(err.Error())
			return "", errors.New(messages.InsertionErr)
		}
	}

	err = transaction.Commit()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.TransactionErr)
	}

	accountBuilder := account.NewBuilder()
	accountBuilder.WithID(accountID)
	for _, role := range userAccount.Roles() {
		accountBuilder.WithRole(role)
	}
	newAccount := accountBuilder.Build()

	signedToken, err := authorization.GenerateTokenFromAccount(*newAccount)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.UnauthorizedErr)
	}

	return signedToken, nil
}

func (repo AccountRepository) SignIn(userAccount account.Account) (string, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	query := `
		select a.id, r.name
		from account a
			inner join account_role ac on ac.account_id = a.id 
			inner join role r on ac.role_id = r.id
		where email = $1 and password = $2;
	`
	rows, err := conn.Queryx(query, userAccount.Email(), userAccount.Password())
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", err
	}
	defer rows.Close()

	accountBuilder := account.NewBuilder()

	for rows.Next() {
		var accountID uuid.UUID
		var roleScan string
		err = rows.Scan(&accountID, &roleScan)
		if err != nil {
			log.Err(err).Msg(err.Error())
			return "", err
		}
		roleBuilder := account.NewRoleBuilder()
		roleBuilder.WithName(roleScan)

		accountBuilder.WithID(accountID)
		accountBuilder.WithRole(*roleBuilder.Build())
	}

	if err = rows.Err(); err != nil {
		log.Err(err).Msg(err.Error())
		return "", err
	}

	signedToken, err := authorization.GenerateTokenFromAccount(*accountBuilder.Build())
	if err != nil {
		log.Err(err).Msg(err.Error())
		return "", errors.New(messages.UnauthorizedErr)
	}

	return signedToken, nil
}

func NewAccountRepository(connector ConnectorManager) *AccountRepository {
	return &AccountRepository{connector}
}
