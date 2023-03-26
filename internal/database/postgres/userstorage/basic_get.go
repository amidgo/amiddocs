package userstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5"
)

const _BASIC_USER_GET = `
	SELECT users.id,users.login,users.name,users.surname,users."fatherName",users.email,users.password,array_agg(roles.role) as userroles 
	FROM users INNER JOIN roles ON users.id = roles."userId" `

func scanUser(q pgx.Row, user *usermodel.UserDTO) error {
	return q.Scan(
		&user.ID,
		&user.Login,
		&user.Name,
		&user.Surname,
		&user.FatherName,
		&user.Email,
		&user.Password,
		&user.Roles,
	)
}

func (u *UserStorage) GetUserRoles(ctx context.Context, userId uint64) ([]userfields.UserRole, *amiderrors.ErrorResponse) {
	roles := make([]userfields.UserRole, 0)
	err := u.p.Pool.QueryRow(ctx,
		`SELECT array_agg(role) from roles inner join users on roles."userId" = users.id where users.id = $1`, userId).Scan(
		&roles,
	)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err)
	}

	return roles, nil
}

func (u *UserStorage) GetUserById(ctx context.Context, userId uint64) (*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	user := new(usermodel.UserDTO)
	q := u.p.Pool.QueryRow(ctx, _BASIC_USER_GET+`WHERE users.id = $1 GROUP BY users.id`, userId)
	err := scanUser(q, user)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err)
	}
	return user, nil
}

func (u *UserStorage) GetUserByLogin(ctx context.Context, login userfields.Login) (*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	user := new(usermodel.UserDTO)
	q := u.p.Pool.QueryRow(ctx, _BASIC_USER_GET+`WHERE users.login = $1 GROUP BY users.id`, login)
	err := scanUser(q, user)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err)
	}
	return user, nil
}

func (u *UserStorage) GetUserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	user := new(usermodel.UserDTO)
	q := u.p.Pool.QueryRow(ctx, _BASIC_USER_GET+`WHERE users.email = $1 GROUP BY users.id`, email)
	err := scanUser(q, user)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err)
	}
	return user, nil
}

func (u *UserStorage) GetAllUsers(ctx context.Context) ([]*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	userList := make([]*usermodel.UserDTO, 0)
	q, err := u.p.Pool.Query(ctx, _BASIC_USER_GET+"GROUP BY users.id")
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err)
	}
	defer q.Close()
	for q.Next() {
		user := new(usermodel.UserDTO)
		err := scanUser(q, user)
		if err != nil {
			return nil, amiderrors.NewInternalErrorResponse(err)
		}
		userList = append(userList, user)
	}
	return userList, nil
}
