package doctypemodel

// create table if not exists document_types(
//     id serial primary key,
//     type varchar(100) not null unique,
//     refresh_time smallint not null,
// );

// create table if not exists document_type_roles (
//     document_type_id int not null references document_types(id),
//     role_id smallint not null references roles(id),
//     primary key (document_type_id, role_id)
// );

const DocTypeTable = "document_types"

type doctype_column string

func (d doctype_column) String() string {
	return string(d)
}

func (d doctype_column) TableName() string {
	return DocTypeTable
}

type doctype_table struct {
	ID          doctype_column
	Type        doctype_column
	RefreshTime doctype_column
}

var SQL = doctype_table{
	ID:          "id",
	Type:        "type",
	RefreshTime: "refresh_time",
}

const DocTypeRoleTable = "document_type_roles"

type doctype_roles_column string

func (d doctype_roles_column) String() string {
	return string(d)
}

func (d doctype_roles_column) TableName() string {
	return DocTypeRoleTable
}

type doctype_roles_table struct {
	DocumentTypeId doctype_roles_column
	RoleId         doctype_roles_column
}

var SQL_ROLES = doctype_roles_table{
	DocumentTypeId: "document_type_id",
	RoleId:         "role_id",
}
