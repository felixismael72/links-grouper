package primary

import (
	"authentication/src/core/domain/account"
)

type AccountManager interface {
	SignUp(userAccount account.Account) (string, error)
	SignIn(userAccount account.Account) (string, error)
}
