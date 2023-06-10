package account

import "github.com/google/uuid"

type RoleBuilder struct {
	role *Role
}

func (builder *RoleBuilder) WithID(id uuid.UUID) *RoleBuilder {
	builder.role.id = id
	return builder
}

func (builder *RoleBuilder) WithName(name string) *RoleBuilder {
	builder.role.name = name
	return builder
}

func (builder *RoleBuilder) Build() *Role {
	return builder.role
}

func NewRoleBuilder() *RoleBuilder {
	return &RoleBuilder{role: &Role{}}
}
