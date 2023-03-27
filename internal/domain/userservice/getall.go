package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserService) GetAllUsers(ctx context.Context) ([]*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	return u.usrrep.GetAllUsers(ctx)
}
