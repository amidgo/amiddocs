package jwttoken

import (
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/golang-jwt/jwt/v4"
)

type amidClaims struct {
	jwt.MapClaims
}

const (
	_USER_ID    = "userid"
	_ROLES      = "roles"
	_EXPIRATION = "exp"
)

func (c *amidClaims) UserID() (uint64, error) {
	return uint64(c.MapClaims[_USER_ID].(float64)), nil
}

func (c *amidClaims) UserRoles() []userfields.Role {
	roleList := make([]userfields.Role, 0)
	for _, r := range c.MapClaims[_ROLES].([]interface{}) {
		roleList = append(roleList, userfields.Role(r.(string)))
	}
	return roleList
}
