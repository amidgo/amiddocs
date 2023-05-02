package reqmodel

import (
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateRequestDTO struct {
	Status       reqfields.Status        `json:"status" db:"status"`
	Count        reqfields.DocumentCount `json:"count" db:"count"`
	DepartmentID uint64                  `json:"departmentId" db:"department_id"`
	DocumentType reqfields.DocumentType  `json:"documentType" db:"document_type"`
	UserID       uint64                  `json:"-"`
}

func (c *CreateRequestDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{c.Count, c.DocumentType}
}

func NewCreateRequest(
	status reqfields.Status,
	count reqfields.DocumentCount,
	departmentId uint64,
	documentType reqfields.DocumentType,
) *CreateRequestDTO {
	return &CreateRequestDTO{
		Status:       status,
		Count:        count,
		DepartmentID: departmentId,
		DocumentType: documentType,
	}
}
