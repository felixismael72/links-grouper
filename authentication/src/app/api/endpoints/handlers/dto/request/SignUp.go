package request

import (
	"authentication/src/core/domain/account"

	"github.com/google/uuid"
)

type SignUp struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	FullName string      `json:"full_name"`
	Password string      `json:"password"`
	RolesIDs []uuid.UUID `json:"roles_ids"`
}

func (dto SignUp) ToDomain() (*account.Account, error) {
	builder := account.NewBuilder()
	builder, err := builder.WithEmail(dto.Email).
		WithUsername(dto.Username).
		WithFullName(dto.FullName).
		WithPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	for _, roleID := range dto.RolesIDs {
		roleBuilder := account.NewRoleBuilder()
		roleBuilder.WithID(roleID)
		builder.WithRole(*roleBuilder.Build())
	}
	return builder.Build(), nil
}
