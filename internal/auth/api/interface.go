package api

import "web-11/internal/auth/provider"

type Usecase interface {
	Authenticate(string, string) (string, error)
	ValidateJWT(string) (*provider.JWTClaims, error)
	Register(string, string) error
}
