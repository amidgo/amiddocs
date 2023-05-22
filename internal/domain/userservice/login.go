package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
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
	refresh_token, err := s.upsertRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	tokenResponse := tokenmodel.NewTokenResponse(token, refresh_token, user.Roles)
	return tokenResponse, nil
}

func (s *userService) upsertRefreshToken(ctx context.Context, userId uint64) (string, error) {
	_, err := s.refreshrep.TokenByUserId(ctx, userId)
	if amiderrors.Is(err, tokenerror.TOKEN_NOT_FOUND) {
		rtoken, err := s.refreshrep.InsertToken(ctx, userId)
		if err != nil {
			return "", err
		}
		return rtoken, nil
	}
	if err != nil {
		return "", err
	}
	rtoken, err := s.refreshrep.UpdateTokenByUserId(ctx, userId)
	if err != nil {
		return "", err
	}
	return rtoken, nil
}
