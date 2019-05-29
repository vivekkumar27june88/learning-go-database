package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:V!VE@!!mysql01@tcp(127.0.0.1:3306)/learning_go")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Database connected")
	}
	defer db.Close()
}
