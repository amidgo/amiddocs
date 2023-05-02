package tokenmodel

import "github.com/amidgo/amiddocs/internal/models/usermodel/userfields"

type TokenResponse struct {
	AccessToken string            `json:"accessToken"`
	Roles       []userfields.Role `json:"roles"`
}

func NewTokenResponse(accessToken string, roles []userfields.Role) *TokenResponse {
	return &TokenResponse{AccessToken: accessToken, Roles: roles}
}
