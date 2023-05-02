package doctempmodel

import (
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type DocumentTemplateDTO struct {
	DepartmentID uint64                 `json:"departmentId"`
	DocumentType reqfields.DocumentType `json:"documentType"`
	Document     []byte                 `json:"document"`
}

func NewCreateDocTemplate(
	depID uint64,
	docType reqfields.DocumentType,
	document []byte,
) *DocumentTemplateDTO {
	return &DocumentTemplateDTO{
		DepartmentID: depID,
		DocumentType: docType,
		Document:     document,
	}
}

func (c *DocumentTemplateDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{c.DocumentType}
}
