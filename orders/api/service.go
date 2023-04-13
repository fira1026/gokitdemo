package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const httpRequestKey = "httpRequest"

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

// CreateOrder creates order in sqlite3
func (s *service) CreateOrder(ctx context.Context, product_name string, quantity int) (int, string, int, error) {
	db, err := sql.Open("sqlite3", "demo_order.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// var order_id int
	// var total_price int

	// Fake log for price
	fmt.Printf("Creating order. Product: %s, price: %d", product_name, 999)

	// Get the headers from the request object in the context
	r, ok := ctx.Value(httpRequestKey).(*http.Request)
	if !ok {
		// TODO: handle error
	}
	email := r.Header.Get("Email")

	// call /v1/products/decrease-quantity to get updated quantity and price
	payload := `{
		"product_name": "` + product_name + `",
		"quantity": ` + strconv.Itoa(quantity) + `
	}`

	req, err := http.Post("http://localhost:8082/v1/products/decrease-quantity", "text/plain", strings.NewReader(payload))
	if err != nil {
		// TODO: handle error
		fmt.Println(err)
	}
	if req.StatusCode != http.StatusOK{
		// TODO: handle error
		fmt.Println(req.StatusCode)
	}
	// TODO: retrive price from req
	// TODO: total_price = quantity * price
	// TODO: insert {email, product_name, quantity and total_price} to orders table

	// return order_id, email, total_price, nil
	return 123, email, 999, nil
}
