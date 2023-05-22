package doctypemodel

import (
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DocumentTypeDTO struct {
	ID          uint64                     `json:"id"`
	Type        doctypefields.DocumentType `json:"type"`
	RefreshTime uint8                      `json:"refresh_time"`
	Roles       []userfields.Role          `json:"role"`
}

func NewDocTypeDTO(id uint64, dtype doctypefields.DocumentType, refreshTime uint8, roles []userfields.Role) *DocumentTypeDTO {
	return &DocumentTypeDTO{ID: id, Type: dtype, RefreshTime: refreshTime, Roles: roles}
}

func (dt *DocumentTypeDTO) ValidatableVariables() []validate.Validatable {
	vList := make([]validate.Validatable, 0)
	vList = append(vList, dt.Type)
	for _, role := range dt.Roles {
		vList = append(vList, role)
	}
	return vList
}
