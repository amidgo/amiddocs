package userservice

import (
	"context"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/user_error_utils"
	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

type loginAmid struct {
	userS *UserService
	err   *amiderrors.ErrorResponse
}

func (a *loginAmid) getUserByLogin(ctx context.Context, login userfields.Login) *usermodel.UserDTO {
	if a.err != nil {
		return nil
	}
	user, err := a.userS.usrrep.GetUserByLogin(ctx, login)
	a.err = err
	return user
}

func (a *loginAmid) verifyPassword(hashPassword userfields.Password, stringPassword userfields.Password) {
	if a.err != nil {
		return
	}
	if !hashPassword.Verify(string(stringPassword)) {
		a.err = usererrorutils.WRONG_PASSWORD
	}
}

func (a *loginAmid) createAccessToken(userId uint64, roles []userfields.UserRole) *tokenmodel.TokenResponse {
	if a.err != nil {
		return nil
	}
	token, err := a.userS.jwtser.CreateUserAccessToken(userId, roles)
	a.err = err
	return tokenmodel.NewTokenResponse(token, roles)
}

func (s *UserService) Login(ctx context.Context, loginform usermodel.LoginForm) (*tokenmodel.TokenResponse, *amiderrors.ErrorResponse) {
	amid := loginAmid{userS: s, err: nil}
	user := amid.getUserByLogin(ctx, loginform.Login)
	amid.verifyPassword(user.Password, loginform.Password)
	tokenResponse := amid.createAccessToken(user.ID, user.Roles)
	if amid.err != nil {
		return nil, amid.err
	}
	return tokenResponse, nil
}
