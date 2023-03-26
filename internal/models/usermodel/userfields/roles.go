package userfields

import (
	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/user_error_utils"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

type UserRole string

const (
	STUDENT   UserRole = "STUDENT"
	ADMIN     UserRole = "ADMIN"
	SECRETARY UserRole = "SECRETARY"
)

func (ur UserRole) Validate() *amiderrors.ErrorResponse {
	for _, r := range []UserRole{STUDENT, ADMIN, SECRETARY} {
		if ur == r {
			return nil
		}
	}
	return usererrorutils.ROLE_NOT_EXIST
}
