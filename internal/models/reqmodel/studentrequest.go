package reqmodel

import (
	"github.com/amidgo/amiddocs/internal/models/doctypemodel"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/pkg/amidtime"
)

type RequestDTO struct {
	ID           uint64                        `json:"id" db:"id"`
	Status       reqfields.Status              `json:"status" db:"status"`
	Count        reqfields.DocumentCount       `json:"count" db:"count"`
	Date         amidtime.DateTime             `json:"date" db:"date"`
	UserID       uint64                        `json:"userId" db:"user_id"`
	DepartmentID uint64                        `json:"departmentId" db:"department_id"`
	DocumentType *doctypemodel.DocumentTypeDTO `json:"documentType" db:"document_type"`
}

func NewRequest(
	id uint64,
	status reqfields.Status,
	count reqfields.DocumentCount,
	date amidtime.DateTime,
	userId uint64,
	departmentId uint64,
	documentType *doctypemodel.DocumentTypeDTO,
) *RequestDTO {
	return &RequestDTO{
		ID:           id,
		Status:       status,
		Count:        count,
		Date:         date,
		UserID:       userId,
		DepartmentID: departmentId,
		DocumentType: documentType,
	}
}
