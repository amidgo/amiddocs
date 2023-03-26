package userrepository

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

type UserRepository interface {
	InsertUser(ctx context.Context, usr *usermodel.UserDTO) (uint64, *amiderrors.ErrorResponse)

	DeleteUser(ctx context.Context, userId uint64) *amiderrors.ErrorResponse

	UpdateName(ctx context.Context, userId uint64, userName string) *amiderrors.ErrorResponse

	UpdateSurname(ctx context.Context, userId uint64, userSurname string) *amiderrors.ErrorResponse

	UpdateFatherName(ctx context.Context, userId uint64, userFatherName string) *amiderrors.ErrorResponse

	UpdateLogin(ctx context.Context, userId uint64, login string) *amiderrors.ErrorResponse

	UpdatePassword(ctx context.Context, userId uint64, hashPassword string) *amiderrors.ErrorResponse

	UpdateEmail(ctx context.Context, userId uint64, email string) *amiderrors.ErrorResponse
}
