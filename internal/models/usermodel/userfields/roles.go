package userfields

import (
	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/usererror"
)

type Role string

const (
	STUDENT   Role = "STUDENT"
	ADMIN     Role = "ADMIN"
	SECRETARY Role = "SECRETARY"
)

func (ur Role) Validate() error {
	for _, r := range []Role{STUDENT, ADMIN, SECRETARY} {
		if ur == r {
			return nil
		}
	}
	return usererrorutils.ROLE_NOT_EXIST
}

func (ur Role) String() string {
	return string(ur)
}
