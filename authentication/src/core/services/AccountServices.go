package services

import (
	"authentication/src/core/domain/account"
	"authentication/src/core/interfaces/primary"
	"authentication/src/core/interfaces/repository"
)

var _ primary.AccountManager = (*AccountServices)(nil)

type AccountServices struct {
	accountRepository repository.AccountLoader
}

func (service AccountServices) SignUp(userAccount account.Account) (string, error) {
	token, err := service.accountRepository.SignUp(userAccount)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service AccountServices) SignIn(userAccount account.Account) (string, error) {
	token, err := service.accountRepository.SignIn(userAccount)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAccountServices(accountRepo repository.AccountLoader) *AccountServices {
	return &AccountServices{accountRepository: accountRepo}
}
