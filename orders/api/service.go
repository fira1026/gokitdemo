package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// to use go-kit the service type is interface
type Service interface {
	CreateOrder(ctx context.Context, product_name string, quantity int) (int, string, int, error)
}

var (
	ErrInvalidProduct = errors.New("invalid product name")
	ErrInvalidQuantiy = errors.New("invalid quantity")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CreateOrder(ctx context.Context, product_name string, quantity int) (int, string, int, error) {
	// create order in sqlite3
	db, err := sql.Open("sqlite3", "demo_order.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

    // var email string
	// var order_id int
	// var total_price int

	// Fake log for price
	fmt.Printf("Creating order. Product: %s, price: %d", product_name, 999)

	// TODO: call /v1/users/validate-token to get user email
	// TODO: call /v1/products/decrease-quantity to get updated quantity and price
	// TODO: total_price = quantity * price
	// TODO: insert {email, product_name, quantity and total_price} to orders table

	// return order_id, email, total_price, nil
	return 123, "fira1026@gmail.com", 999, nil
}
