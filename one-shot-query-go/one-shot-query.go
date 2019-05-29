package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:V!VE@!!mysql01@(127.0.0.1:3306)/learning_go")
	if err == nil {
		log.Println("MySQL Connection Established")
	} else {
		log.Fatalln("Error in Establishing connection with mysql")
	}
	defer db.Close()

	var str1, str2, str3 string
	err = db.QueryRow("SELECT * FROM programming_languages WHERE id = 2").Scan(&str1, &str2, &str3)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	log.Println(str1 + str2 + str3)
}
