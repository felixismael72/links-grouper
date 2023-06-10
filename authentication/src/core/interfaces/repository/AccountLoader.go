package repository

import (
	"authentication/src/core/domain/account"
)

type AccountLoader interface {
	SignUp(userAccount account.Account) (string, error)
	SignIn(userAccount account.Account) (string, error)
}
