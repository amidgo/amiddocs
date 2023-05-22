package depmodel

import "github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"

type DepartmentTypes struct {
	Department *DepartmentDTO               `json:"department"`
	Types      []doctypefields.DocumentType `json:"types"`
}

func NewDepTypes(dep *DepartmentDTO, types []doctypefields.DocumentType) *DepartmentTypes {
	return &DepartmentTypes{Department: dep, Types: types}
}
