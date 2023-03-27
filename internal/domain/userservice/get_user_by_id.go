package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserService) GetUserById(ctx context.Context, id uint64) (*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	return u.usrrep.GetUserById(ctx, id)
}
