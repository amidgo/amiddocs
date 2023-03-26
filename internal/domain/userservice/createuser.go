package userservice

import (
	"context"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/user_error_utils"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

type createUserAmid struct {
	userS *UserService
	err   *amiderrors.ErrorResponse
}

func (a *createUserAmid) checkEmail(ctx context.Context, email userfields.Email) {
	if a.err != nil {
		return
	}
	usr, _ := a.userS.usrrep.GetUserByEmail(ctx, email)
	if usr != nil {
		a.err = usererrorutils.EMAIL_ALREADY_EXIST
		return
	}
}

func (a *createUserAmid) generateLoginAndPassword(u *usermodel.CreateUserDTO) (userfields.Login, userfields.Password) {
	if a.err != nil {
		return "", ""
	}
	login, password, err := u.GenerateLoginAndPassword()
	a.err = err
	return login, password
}

func (a *createUserAmid) insertUser(ctx context.Context, user *usermodel.UserDTO) uint64 {
	if a.err != nil {
		return 0
	}
	userId, err := a.userS.usrrep.InsertUser(ctx, user)
	a.err = err
	return userId
}

func (s *UserService) CreateUser(ctx context.Context, u *usermodel.CreateUserDTO) (*usermodel.UserDTO, *amiderrors.ErrorResponse) {
	amid := createUserAmid{userS: s, err: nil}
	amid.checkEmail(ctx, u.Email)
	login, password := amid.generateLoginAndPassword(u)
	user := usermodel.NewUserDTO(0, login, password, u.Name, u.Surname, u.FatherName, u.Email, u.Roles)
	userId := amid.insertUser(ctx, user)
	user.ID = userId
	if amid.err != nil {
		return nil, amid.err
	}
	return user, nil
}
