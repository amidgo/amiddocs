package userstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func basicUserQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT 
		users.id,users.login,users.name,users.surname,users.father_name,
		users.email,users.password,array_agg(roles.role) as userroles 
	FROM users 
		INNER JOIN user_roles ON users.id = user_roles.user_id INNER JOIN roles ON roles.id = user_roles.role_id
		%s
		GROUP BY users.id
		`,
		query,
	)
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func scanUser(q Scanner, user *usermodel.UserDTO) error {
	err := q.Scan(
		&user.ID,
		&user.Login,
		&user.Name,
		&user.Surname,
		&user.FatherName,
		&user.Email,
		&user.Password,
		&user.Roles,
	)
	return err
}

func (u *userStorage) getUserByQuery(ctx context.Context, query string, args ...interface{}) (*usermodel.UserDTO, error) {
	user := new(usermodel.UserDTO)
	q := u.p.Pool.QueryRow(ctx, query, args...)
	err := scanUser(q, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userStorage) UserById(ctx context.Context, userId uint64) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE users.id = $1`), userId)
	if err != nil {
		return nil, userError(err, amiderrors.NewCause("user by id query", "UserById", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) UserByLogin(ctx context.Context, login userfields.Login) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE users.login = $1`), login)
	if err != nil {
		return nil, userError(err, amiderrors.NewCause("user by login query", "UserByLogin", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE users.email = $1`), email)
	if err != nil {
		return nil, userError(err, amiderrors.NewCause("user by email query", "UserByEmail", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) AllUsers(ctx context.Context) ([]*usermodel.UserDTO, error) {
	userList := make([]*usermodel.UserDTO, 0)
	q, err := u.p.Pool.Query(ctx, basicUserQuery(""))
	if err != nil {
		return nil, userError(err, amiderrors.NewCause("query all user", "AllUsers", _PROVIDER))
	}
	defer q.Close()
	for q.Next() {
		user := new(usermodel.UserDTO)
		err := scanUser(q, user)
		if err != nil {
			return nil, userError(err, amiderrors.NewCause(
				fmt.Sprintf("scan user number %d", len(userList)),
				"AllUsers",
				_PROVIDER,
			))
		}
		userList = append(userList, user)
	}
	return userList, nil
}
