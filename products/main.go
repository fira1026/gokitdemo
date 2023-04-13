package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"products/api"

	"github.com/go-kit/kit/log"
	_ "github.com/mattn/go-sqlite3"
)

// main is the uses entry point
func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8081", "caller", log.DefaultCaller)

	initDb()

	// Start products api server
	r := api.NewHttpServer(api.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", r))
}

func initDb() {
	db, err := sql.Open("sqlite3", "demo_product.db")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, table_check := db.Query("SELECT * FROM products")
	if table_check != nil {
		sts := `
		DROP TABLE IF EXISTS products;
		CREATE TABLE products(id INTEGER PRIMARY KEY, name TEXT, price INT, quantity INT);
		INSERT INTO products(name, price, quantity) VALUES('Apple', '20', '50');
		INSERT INTO products(name, price, quantity) VALUES('banana', '10', '100');
		`
		_, err = db.Exec(sts)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("table products created")
	}

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
		return
    }

	defer rows.Close()

    for rows.Next() {

        var id int
        var name string
        var price int
		var quantity int

        err = rows.Scan(&id, &name, &price, &quantity)

        if err != nil {
			fmt.Println(err)
			return
        }

        fmt.Printf("%d %s %d %d\n" , id, name, price, quantity)
	}
}
