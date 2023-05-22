package depmodel

import (
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DepartmentDTO struct {
	ID        uint64              `json:"id" db:"id"`
	Name      depfields.Name      `json:"name" db:"name"`
	ShortName depfields.ShortName `json:"-" db:"short_name"`
}

func NewDepartmentDTO(id uint64, name depfields.Name, shortName depfields.ShortName) *DepartmentDTO {
	return &DepartmentDTO{ID: id, Name: name, ShortName: shortName}
}

func (d *DepartmentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{d.Name, d.ShortName}
}

type CreateDepartmentDTO struct {
	Name      depfields.Name      `json:"name"`
	ShortName depfields.ShortName `json:"shortName"`
}

func (cd *CreateDepartmentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{cd.Name, cd.ShortName}
}
