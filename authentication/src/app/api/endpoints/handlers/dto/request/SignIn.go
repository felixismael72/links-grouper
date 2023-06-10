package request

import "authentication/src/core/domain/account"

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto SignIn) ToDomain() (*account.Account, error) {
	builder, err := account.NewBuilder().
		WithEmail(dto.Email).
		WithPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	return builder.Build(), nil
}
