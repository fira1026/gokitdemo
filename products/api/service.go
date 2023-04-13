package api

import (
	"database/sql"
	"fmt"
	"context"
	"errors"

	_ "github.com/mattn/go-sqlite3"
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
	// decrease product quantity in sqlite3
	db, err := sql.Open("sqlite3", "demo_product.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT quantity FROM products WHERE name=$1;`
	var qt int

	row := db.QueryRow(sqlStatement, product_name)
	switch err := row.Scan(&qt); err {
	case sql.ErrNoRows:
	    fmt.Println("No rows were returned!")
	case nil:
	    fmt.Println(product_name)
	default:
	panic(err)
	}

	if qt < quantity {
		fmt.Println("stock quantity is not enought")
		return -1, ErrInvalidQuantiy
	}

	new_quantity := qt - quantity

	db.Exec("UPDATE products SET quantity = ? WHERE name = ?", new_quantity, product_name)
	fmt.Printf("UPDATE %s quantity to %d\n", product_name, new_quantity)
	return new_quantity, nil
}
