package account

import (
	"github.com/google/uuid"
)

type Account struct {
	id       uuid.UUID
	email    string
	username string
	fullName string
	password string
	hash     string

	roles []Role
}

func (a Account) ID() uuid.UUID {
	return a.id
}

func (a Account) Email() string {
	return a.email
}

func (a Account) Username() string {
	return a.username
}

func (a Account) FullName() string {
	return a.fullName
}

func (a Account) Password() string {
	return a.password
}

func (a Account) Hash() string {
	return a.hash
}

func (a Account) Roles() []Role {
	return a.roles
}
