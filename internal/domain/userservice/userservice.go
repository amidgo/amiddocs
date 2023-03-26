package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

type UserRepositoryInterface interface {
	InsertUser(ctx context.Context, user *usermodel.UserDTO) (uint64, *amiderrors.ErrorResponse)
	UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) *amiderrors.ErrorResponse
	UpdateName(ctx context.Context, userId uint64, name userfields.Name) *amiderrors.ErrorResponse
	UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) *amiderrors.ErrorResponse
	UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) *amiderrors.ErrorResponse
	UpdatePassword(ctx context.Context, userId uint64, password string) *amiderrors.ErrorResponse
	UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) *amiderrors.ErrorResponse
	GetUserById(ctx context.Context, userId uint64) (*usermodel.UserDTO, *amiderrors.ErrorResponse)
	GetUserByLogin(ctx context.Context, login userfields.Login) (*usermodel.UserDTO, *amiderrors.ErrorResponse)
	GetUserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, *amiderrors.ErrorResponse)
	GetAllUsers(ctx context.Context) ([]*usermodel.UserDTO, *amiderrors.ErrorResponse)
}

type JWTService interface {
	CreateUserAccessToken(userId uint64, roles []userfields.UserRole) (string, *amiderrors.ErrorResponse)
}
type UserService struct {
	usrrep UserRepositoryInterface
	jwtser JWTService
}

func New(userrep UserRepositoryInterface, jwtser JWTService) *UserService {
	return &UserService{usrrep: userrep, jwtser: jwtser}
}
