package config

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB database
var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@/golangcms?parseTime=true")
	if err != nil {
		panic(err)
	}

	DB = db
	fmt.Println("You connected to your mysql database.")
}
