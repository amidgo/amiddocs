package jwttoken

import (
	tokenerrorutils "github.com/amidgo/amiddocs/internal/errorutils/token_error_utils"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/golang-jwt/jwt/v4"
)

type amidClaims struct {
	jwt.MapClaims
}

const (
	_USER_ID    = "userid"
	_STUDENT_ID = "studentid"
	_ROLES      = "roles"
	_EXPIRATION = "exp"
)

func (c *amidClaims) GetUserId() (uint64, *amiderrors.ErrorResponse) {
	return c.MapClaims[_USER_ID].(uint64), nil
}

func (c *amidClaims) GetStudentId() (uint64, *amiderrors.ErrorResponse) {
	if c.MapClaims[_ROLES].(userfields.UserRole) != userfields.STUDENT {
		return 0, tokenerrorutils.FORBIDDEN
	}
	return c.MapClaims[_STUDENT_ID].(uint64), nil
}

func (c *amidClaims) GetUserRoles() []userfields.UserRole {
	roleList := make([]userfields.UserRole, 0)
	for _, r := range c.MapClaims[_ROLES].([]interface{}) {
		roleList = append(roleList, userfields.UserRole(r.(string)))
	}
	return roleList
}
