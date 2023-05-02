package doctypemodel

import (
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DocumentTypeDTO struct {
	ID          uint                   `json:"id"`
	Type        reqfields.DocumentType `json:"type"`
	RefreshTime uint8                  `json:"refresh_time"`
	Role        userfields.Role        `json:"role"`
}

func NewDocTypeDTO(id uint, dtype reqfields.DocumentType, refreshTime uint8, role userfields.Role) *DocumentTypeDTO {
	return &DocumentTypeDTO{ID: id, Type: dtype, RefreshTime: refreshTime, Role: role}
}

func (dt *DocumentTypeDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{dt.Role, dt.Type}
}
