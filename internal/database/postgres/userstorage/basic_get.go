package userstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/sqlutils"
)

var (
	userRolesQuery = fmt.Sprintf(
		`
		SELECT array_agg(%s) 
		FROM %s 
		INNER JOIN %s ON %s = %s
		WHERE %s = $1
		GROUP BY %s
		`,
		// select roles
		sqlutils.Full(usermodel.SQL_ROLES.Role),

		// roles table
		usermodel.RolesTable,

		// inner join roles on user roles by role id
		usermodel.UserRolesTable,
		sqlutils.Full(usermodel.SQL_ROLES.ID),
		sqlutils.Full(usermodel.SQL_USER_ROLES.RoleId),

		// where user roles userid
		sqlutils.Full(usermodel.SQL_USER_ROLES.UserId),

		// group by columns
		sqlutils.Full(usermodel.SQL_USER_ROLES.UserId),
	)
)

func basicUserQuery(query string) string {
	return fmt.Sprintf(
		`
	SELECT 
		%s,%s,%s,%s,%s,%s,%s,array_agg(%s) as userroles 
	FROM %s 
		INNER JOIN %s ON %s = %s 
		INNER JOIN %s ON %s = %s
		%s
		GROUP BY %s
		`,
		// selectable variables
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(usermodel.SQL.Login),
		sqlutils.Full(usermodel.SQL.Name),
		sqlutils.Full(usermodel.SQL.Surname),
		sqlutils.Full(usermodel.SQL.FatherName),
		sqlutils.Full(usermodel.SQL.Email),
		sqlutils.Full(usermodel.SQL.Password),
		sqlutils.Full(usermodel.SQL_ROLES.Role),

		// user table
		usermodel.UserTable,

		// inner join user table on user roles table by user id
		usermodel.UserRolesTable,
		sqlutils.Full(usermodel.SQL.ID),
		sqlutils.Full(usermodel.SQL_USER_ROLES.UserId),

		// inner join roles on user roles by role id
		usermodel.RolesTable,
		sqlutils.Full(usermodel.SQL_ROLES.ID),
		sqlutils.Full(usermodel.SQL_USER_ROLES.RoleId),

		query,

		// group by column
		sqlutils.Full(usermodel.SQL.ID),
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

func (u *userStorage) UserRoles(ctx context.Context, userId uint64) ([]userfields.Role, error) {
	roles := make([]userfields.Role, 0)
	fmt.Println(userRolesQuery)
	err := u.p.Pool.QueryRow(ctx,
		userRolesQuery,
		userId,
	).Scan(&roles)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("user roles query row", "UserRoles", _PROVIDER))
	}
	return roles, nil
}

func (u *userStorage) UserById(ctx context.Context, userId uint64) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE `+sqlutils.Full(usermodel.SQL.ID)+` = $1`), userId)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("user by id query", "UserById", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) UserByLogin(ctx context.Context, login userfields.Login) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE `+sqlutils.Full(usermodel.SQL.Login)+` = $1`), login)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("user by login query", "UserByLogin", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error) {
	usr, err := u.getUserByQuery(ctx, basicUserQuery(`WHERE `+sqlutils.Full(usermodel.SQL.Email)+` = $1`), email)
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("user by email query", "UserByEmail", _PROVIDER))
	}
	return usr, nil
}

func (u *userStorage) AllUsers(ctx context.Context) ([]*usermodel.UserDTO, error) {
	userList := make([]*usermodel.UserDTO, 0)
	q, err := u.p.Pool.Query(ctx, basicUserQuery(""))
	if err != nil {
		return nil, UserError(err, amiderrors.NewCause("query all user", "AllUsers", _PROVIDER))
	}
	defer q.Close()
	for q.Next() {
		user := new(usermodel.UserDTO)
		err := scanUser(q, user)
		if err != nil {
			return nil, UserError(err, amiderrors.NewCause(
				fmt.Sprintf("scan user number %d", len(userList)),
				"AllUsers",
				_PROVIDER,
			))
		}
		userList = append(userList, user)
	}
	return userList, nil
}
