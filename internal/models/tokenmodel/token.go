package tokenmodel

import "github.com/amidgo/amiddocs/internal/models/usermodel/userfields"

type TokenResponse struct {
	AccessToken string                `json:"accessToken"`
	Roles       []userfields.UserRole `json:"roles"`
}

func NewTokenResponse(accessToken string, roles []userfields.UserRole) *TokenResponse {
	return &TokenResponse{AccessToken: accessToken, Roles: roles}
}
