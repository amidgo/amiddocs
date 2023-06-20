package stdocmodel

const StudentDocumentTable = "student_documents"

const (
	StudentDocumentsNumberUniqueConstraint    = "student_documents_number_unique"
	StudentDocumentsStudentIdUniqueConstraint = "student_documents_student_id_unique"

	ForeignKey_StudentDocuments__Students = "fk_student_documents__students"
)

type student_document_column string

func (st student_document_column) String() string {
	return string(st)
}

func (st student_document_column) TableName() string {
	return StudentDocumentTable
}

var SQL = struct {
	ID                 student_document_column
	StudentId          student_document_column
	DocNumber          student_document_column
	OrderNumber        student_document_column
	OrderDate          student_document_column
	EducationStartDate student_document_column
}{
	ID:                 "id",
	StudentId:          "student_id",
	DocNumber:          "doc_number",
	OrderNumber:        "order_number",
	OrderDate:          "order_date",
	EducationStartDate: "education_start_date",
}
