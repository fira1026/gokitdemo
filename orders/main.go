package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"orders/api"

	"github.com/go-kit/kit/log"
	_ "github.com/mattn/go-sqlite3"
)

// main is the uses entry point
func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", "8083", "caller", log.DefaultCaller)

	initDb()

	// Start products api server
	r := api.NewHttpServer(api.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", "8083")
	logger.Log("err", http.ListenAndServe(":8083", r))
}

func initDb() {
	db, err := sql.Open("sqlite3", "demo_order.db")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, table_check := db.Query("SELECT * FROM orders")
	if table_check != nil {
		sts := `
		CREATE TABLE orders(id INTEGER PRIMARY KEY, email TEXT, product_name TEXT, quantity INT, total_price INT);
		`
		_, err = db.Exec(sts)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("table orders created")
	}

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		fmt.Println(err)
		return
    }

	defer rows.Close()

    for rows.Next() {

        var id int
		var email string
        var product_name string
		var quantity int
        var total_price int

        err = rows.Scan(&id, &email, &product_name, &quantity, &total_price)

        if err != nil {
			fmt.Println(err)
			return
        }

        fmt.Printf("%d %s %s %d %d\n" , id, email, product_name, quantity, total_price)
	}
}
