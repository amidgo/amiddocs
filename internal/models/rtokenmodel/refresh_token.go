package rtokenmodel

import (
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/google/uuid"
)

type RefreshToken struct {
	UserId  uint64
	Expired amidtime.DateTime
	Token   uuid.UUID
}
