package doctempmodel

import (
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type CreateTemplateDTO struct {
	DepartmentID uint64                     `json:"departmentId"`
	DocumentType doctypefields.DocumentType `json:"documentType"`
	Document     []byte                     `json:"document"`
}

func NewCreateDocTemplate(
	depID uint64,
	docType doctypefields.DocumentType,
	document []byte,
) *CreateTemplateDTO {
	return &CreateTemplateDTO{
		DepartmentID: depID,
		DocumentType: docType,
		Document:     document,
	}
}

func (c *CreateTemplateDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{c.DocumentType}
}

type DocumentTemplateDTO struct {
	DepartmentID   uint64 `json:"departmentId"`
	DocumentTypeID uint64 `json:"documentType"`
	Document       []byte `json:"document"`
}

func NewDocumentTemplateDTO(depId, documentTypeId uint64, document []byte) *DocumentTemplateDTO {
	return &DocumentTemplateDTO{DepartmentID: depId, DocumentTypeID: documentTypeId, Document: document}
}
