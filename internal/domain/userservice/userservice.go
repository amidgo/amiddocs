package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

type userRepositoryInterface interface {
	InsertUser(ctx context.Context, user *usermodel.UserDTO) (*usermodel.UserDTO, error)
	UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) error
	UpdateName(ctx context.Context, userId uint64, name userfields.Name) error
	UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) error
	UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) error
	UpdatePassword(ctx context.Context, userId uint64, password string) error
}

type userProvider interface {
	UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) error
	UserById(ctx context.Context, userId uint64) (*usermodel.UserDTO, error)
	UserByLogin(ctx context.Context, login userfields.Login) (*usermodel.UserDTO, error)
	UserByEmail(ctx context.Context, email userfields.Email) (*usermodel.UserDTO, error)
	AllUsers(ctx context.Context) ([]*usermodel.UserDTO, error)
}
type jwtService interface {
	CreateAccessToken(userId uint64, roles []userfields.Role) (string, error)
}

type encrypter interface {
	Hash(input string) (string, error)
	Verify(hashPassword string, password string) bool
}
type userService struct {
	encrypter encrypter
	usrrep    userRepositoryInterface
	userprov  userProvider
	jwtser    jwtService
}

func New(userrep userRepositoryInterface, jwtser jwtService, usrProv userProvider, encrypter encrypter) *userService {
	return &userService{usrrep: userrep, userprov: usrProv, jwtser: jwtser, encrypter: encrypter}
}
