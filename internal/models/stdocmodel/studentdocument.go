package stdocmodel

import (
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/validate"
)

type StudentDocumentDTO struct {
	ID                 uint64                  `json:"id" db:"id"`
	StudentId          uint64                  `json:"studentId"`
	DocNumber          stdocfields.DocNumber   `json:"docNumber" db:"doc_number"`
	OrderNumber        stdocfields.OrderNumber `json:"orderNumber" db:"order_number"`
	OrderDate          amidtime.Date           `json:"orderDate" db:"order_date"`
	EducationStartDate amidtime.Date           `json:"studyStartDate" db:"education_start_date"`
}

func (sd *StudentDocumentDTO) ValidatableVariables() []validate.Validatable {
	return []validate.Validatable{sd.DocNumber, sd.OrderNumber}
}

func NewStudentDocumentDTO(
	id uint64,
	studentId uint64,
	docNumber stdocfields.DocNumber,
	orderNumber stdocfields.OrderNumber,
	orderDate amidtime.Date,
	studyStartDate amidtime.Date,
) *StudentDocumentDTO {
	return &StudentDocumentDTO{ID: id, StudentId: studentId, DocNumber: docNumber, OrderNumber: orderNumber, OrderDate: orderDate, EducationStartDate: studyStartDate}
}
