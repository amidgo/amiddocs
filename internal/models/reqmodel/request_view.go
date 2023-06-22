package reqmodel

import (
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/reqmodel/reqfields"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

type fio struct {
	Name       userfields.Name       `json:"name"`
	Surname    userfields.Surname    `json:"surname"`
	FatherName userfields.FatherName `json:"fatherName"`
}

type RequestViewDTO struct {
	ID            uint64                     `json:"id"`
	FIO           fio                        `json:"fio"`
	Status        reqfields.Status           `json:"status"`
	DocumentType  doctypefields.DocumentType `json:"documentType"`
	DocumentCount reqfields.DocumentCount    `json:"documentCount"`
}

func NewRequestViewDTO(
	id uint64,
	name userfields.Name,
	surname userfields.Surname,
	fatherName userfields.FatherName,
	status reqfields.Status,
	docType doctypefields.DocumentType,
	documentCount reqfields.DocumentCount,
) *RequestViewDTO {
	return &RequestViewDTO{
		ID: id,
		FIO: fio{
			Name:       name,
			Surname:    surname,
			FatherName: fatherName,
		},
		Status:        status,
		DocumentType:  docType,
		DocumentCount: documentCount,
	}
}
