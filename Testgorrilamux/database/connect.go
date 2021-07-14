package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cosmetics_store")
	if err != nil {
		log.Printf("Error %s when opening DBn", err)
		return nil, err
	}
	DB = db

	return db, nil
}
