package api

import (
	"context"
	"errors"
)

// to use go-kit the service type is interface
type Service interface {
	 DecreaseProductQuantity(ctx context.Context, product_name string, quantity int) (int, error)
}

var (
	ErrInvalidProduct  = errors.New("invalid product name")
	ErrInvalidQuantiy = errors.New("invalid quantity")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) DecreaseProductQuantity(ctx context.Context, product_name string, quantity int) (int, error) {
	var q = 100

	// TODO: decrease product quantity in sqlite3

	return q, nil
}
