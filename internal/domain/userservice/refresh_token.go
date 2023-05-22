package userservice

import (
	"context"
	"time"

	"github.com/amidgo/amiddocs/internal/errorutils/tokenerror"
	"github.com/amidgo/amiddocs/internal/models/tokenmodel"
)

func (s *userService) RefreshToken(ctx context.Context, oldRefreshToken string, userId uint64) (*tokenmodel.TokenResponse, error) {
	roles, err := s.userprov.UserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}
	token, err := s.jwtser.CreateAccessToken(userId, roles)
	if err != nil {
		return nil, err
	}
	refresh_token, err := s.refreshToken(ctx, oldRefreshToken, userId)
	if err != nil {
		return nil, err
	}
	return tokenmodel.NewTokenResponse(token, refresh_token, roles), nil
}

func (s *userService) refreshToken(ctx context.Context, oldRefreshToken string, userId uint64) (string, error) {
	token, err := s.refreshrep.Token(ctx, userId, oldRefreshToken)
	if err != nil {
		return "", err
	}
	if token.Expired.T().Before(time.Now()) {
		return "", tokenerror.REFRESH_TOKEN_EXPIRED
	}
	rtoken, err := s.refreshrep.UpdateTokenByUserId(ctx, userId)
	if err != nil {
		return "", err
	}
	return rtoken, nil
}
