package usecase

import "web-11/internal/auth/provider"

type Provider interface {
	CreateUser(string, string) error
	CheckUserByUsername(string) (bool, error)
	CheckPassword(string, string) (bool, error)
}

type JWTProvider interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*provider.JWTClaims, error)
}
