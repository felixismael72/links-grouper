package account

import (
	"authentication/src/utils/encrypt"

	"github.com/google/uuid"
)

type Builder struct {
	account *Account
}

func (builder *Builder) WithID(id uuid.UUID) *Builder {
	builder.account.id = id
	return builder
}

func (builder *Builder) WithEmail(email string) *Builder {
	builder.account.email = email
	return builder
}

func (builder *Builder) WithUsername(username string) *Builder {
	builder.account.username = username
	return builder
}

func (builder *Builder) WithFullName(fullName string) *Builder {
	builder.account.fullName = fullName
	return builder
}

func (builder *Builder) WithPassword(password string) (*Builder, error) {
	hashedPassword, hash, err := encrypt.HashPassword(password)
	if err != nil {
		return builder, err
	}

	builder.account.password = hashedPassword
	builder.account.hash = hash
	return builder, nil
}

func (builder *Builder) WithRole(role Role) *Builder {
	builder.account.roles = append(builder.account.roles, role)
	return builder
}

func (builder *Builder) Build() *Account {
	return builder.account
}

func NewBuilder() *Builder {
	return &Builder{account: &Account{}}
}
