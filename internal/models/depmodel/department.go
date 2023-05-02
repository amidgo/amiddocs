package depmodel

import (
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DepartmentDTO struct {
	ID        uint64              `json:"id" db:"id"`
	Name      depfields.Name      `json:"name" db:"name"`
	ShortName depfields.ShortName `json:"shortName" db:"short_name"`
}

func (d *DepartmentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{d.Name, d.ShortName}
}
