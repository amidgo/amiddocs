package reqmodel

const RequestTable = "requests"

const (
	ForeignKey_Requests__RequestStatus = "fk_requests__request_status"
	ForeignKey_Requests__Users         = "fk_requests__users"
	ForeignKey_Requests__Departments   = "fk_requests__departments"
	ForeignKey_Requests__DocumentTypes = "fk_requests__document_types"
)

type request_column string

func (r request_column) String() string {
	return string(r)
}

func (r request_column) TableName() string {
	return RequestTable
}

type request_table struct {
	ID             request_column
	StatusId       request_column
	Count          request_column
	Date           request_column
	UserID         request_column
	DepartmentID   request_column
	DocumentTypeId request_column
}

var SQL = request_table{
	ID:             "id",
	StatusId:       "status_id",
	Count:          "count",
	Date:           "date",
	UserID:         "user_id",
	DepartmentID:   "department_id",
	DocumentTypeId: "document_type_id",
}

const RequestStatusTable = "request_status"

type status_column string

func (s status_column) String() string {
	return string(s)
}

func (s status_column) TableName() string {
	return RequestStatusTable
}

type status_table struct {
	ID     status_column
	Status status_column
}

var SQL_STATUS = status_table{
	ID:     "id",
	Status: "status",
}
