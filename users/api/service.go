package api

import (
	auth "users/auth"
	"context"
	"errors"
)

// to use go-kit the service type is interface
type Service interface {
	 ValidateUser(ctx context.Context, mail, password string) (string, error)
	 ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) ValidateUser(ctx context.Context, email, password string) (string, error) {
	//@TODO add register function and use database to store user credentials
	if !(email == "fira1026@gmail.com" && password == "1234567") {
		return "", ErrInvalidUser
	}
	token, err := auth.NewToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := auth.ParseToken(token)
	if err != nil {
		return "", ErrInvalidToken
	}
	tData, err := auth.GetClaims(t)
	if err != nil {
		return "", ErrInvalidToken
	}
	return tData["email"].(string), nil
}
