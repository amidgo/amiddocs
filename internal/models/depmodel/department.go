package depmodel

import (
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DepartmentDTO struct {
	ID        uint64              `json:"id" db:"id"`
	ImageUrl  depfields.ImageUrl  `json:"imageUrl"`
	Name      depfields.Name      `json:"name" db:"name"`
	ShortName depfields.ShortName `json:"-" db:"short_name"`
}
type StudyDepartmentDTO struct {
	StudyDepartmentID uint64 `json:"id"`
	DepartmentDTO
}

func NewStudyDepartmentDTO(studyDepartmentId uint64, departmentDTO DepartmentDTO) *StudyDepartmentDTO {
	return &StudyDepartmentDTO{StudyDepartmentID: studyDepartmentId, DepartmentDTO: departmentDTO}
}

func NewDepartmentDTO(id uint64, name depfields.Name, shortName depfields.ShortName) *DepartmentDTO {
	return &DepartmentDTO{ID: id, Name: name, ShortName: shortName}
}

func (d *DepartmentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{d.Name, d.ShortName}
}

type CreateDepartmentDTO struct {
	Name      depfields.Name      `json:"name" csv:"Название"`
	ShortName depfields.ShortName `json:"shortName" csv:"Инициалы"`
}

func (cd *CreateDepartmentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{cd.Name, cd.ShortName}
}
