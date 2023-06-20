package doctempmodel

const DocTempTable = "document_templates"

const (
	ForeignKey_DocumentTemplates__Departments   = "fk_document_templates__departments"
	ForeignKey_DocumentTemplates__DocumentTypes = "fk_document_templates__document_types"
)

type doctemp_column string

func (d doctemp_column) String() string {
	return string(d)
}

func (d doctemp_column) TableName() string {
	return DocTempTable
}

type doctemp_table struct {
	DepartmentId   doctemp_column
	DocumentTypeId doctemp_column
	Data           doctemp_column
}

var SQL = doctemp_table{
	DepartmentId:   "department_id",
	DocumentTypeId: "document_type_id",
	Data:           "data",
}
