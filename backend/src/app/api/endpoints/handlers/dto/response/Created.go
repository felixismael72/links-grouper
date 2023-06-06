package response

import "github.com/google/uuid"

type Created struct {
	ID uuid.UUID `json:"id"`
}
