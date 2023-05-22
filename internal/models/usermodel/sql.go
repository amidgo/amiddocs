package usermodel

const UserTable = "users"

type user_column string

func (u user_column) TableName() string {
	return "users"
}
func (u user_column) String() string {
	return string(u)
}

type user_sql struct {
	ID         user_column
	Login      user_column
	Password   user_column
	Name       user_column
	Surname    user_column
	FatherName user_column
	Email      user_column
}

var SQL = user_sql{
	ID:         "id",
	Login:      "login",
	Password:   "password",
	Name:       "name",
	Surname:    "surname",
	FatherName: "father_name",
	Email:      "email",
}

const RolesTable = "roles"

type roles_column string

func (r roles_column) String() string {
	return string(r)
}

func (u roles_column) TableName() string {
	return RolesTable
}

type sql_roles struct {
	ID   roles_column
	Role roles_column
}

var SQL_ROLES = sql_roles{
	ID:   "id",
	Role: "role",
}

const UserRolesTable = "user_roles"

type user_roles_column string

func (ur user_roles_column) String() string {
	return string(ur)
}
func (ur user_roles_column) TableName() string {
	return UserRolesTable
}

var SQL_USER_ROLES = struct {
	TABLE_NAME user_roles_column
	UserId     user_roles_column
	RoleId     user_roles_column
}{
	TABLE_NAME: "user_roles",
	UserId:     "user_id",
	RoleId:     "role_id",
}
