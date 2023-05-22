package tokenmodel

import "github.com/amidgo/amiddocs/internal/models/usermodel/userfields"

type TokenResponse struct {
	AccessToken  string            `json:"accessToken"`
	RefreshToken string            `json:"refreshToken"`
	Roles        []userfields.Role `json:"roles"`
}

func NewTokenResponse(accessToken, refreshToken string, roles []userfields.Role) *TokenResponse {
	return &TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken, Roles: roles}
}
