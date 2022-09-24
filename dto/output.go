package dto

import "github.com/ochom/sdp-lib/models"

// AuthResponse ...
type AuthResponse struct {
	User         models.User         `json:"user"`
	Organization models.Organization `json:"organization"`
	AccessToken  string              `json:"accessToken"`
	RefreshToken string              `json:"refreshToken"`
}
