package rtokenmodel

import "github.com/google/uuid"

type RefreshDTO struct {
	UserId uint64    `json:"userId"`
	Token  uuid.UUID `json:"token"`
}
