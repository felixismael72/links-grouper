package account

import (
	"strings"

	"github.com/google/uuid"
)

type Role struct {
	id   uuid.UUID
	name string
}

func (r Role) ID() uuid.UUID {
	return r.id
}

func (r Role) Name() string {
	return strings.ToUpper(r.name)
}

func (r Role) IsZero() bool {
	return r == Role{}
}
