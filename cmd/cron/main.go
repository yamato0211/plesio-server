package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@protocol(%s:%s)/%s",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"),
	))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQLクエリの実行
	_, err = db.Exec("UPDATE users SET IsLogined = false WHERE IsLogined = true")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update successful")
}
