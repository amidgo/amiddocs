package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
)

func (s *userService) Login(
	ctx context.Context,
	loginform *usermodel.LoginForm,
) (*tokenmodel.TokenResponse, error) {

	user, err := s.userprov.UserByLogin(ctx, loginform.Login)
	if err != nil {
		return nil, err
	}
	if !s.encrypter.Verify(string(user.Password), string(loginform.Password)) {
		return nil, usererror.WRONG_PASSWORD
	}
	token, err := s.jwtser.CreateAccessToken(user.ID, user.Roles)
	if err != nil {
		return nil, err
	}
	tokenRespone := tokenmodel.NewTokenResponse(token, user.Roles)
	return tokenRespone, nil
}
