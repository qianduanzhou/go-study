package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func main() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	fmt.Println(database, err)
}
