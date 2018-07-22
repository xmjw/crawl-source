package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var connection = DatabaseConnection()

func WriteDomain(domain string) {
	_, err := connection.Exec(
		"INSERT INTO sources (domain) VALUES ($1)",
		domain)

	if err != nil {
		panic(err)
	}
}

func DatabaseConnection() *sql.DB {
	connectionString := os.Getenv("SOURCE_URL_DATABASE")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	return db
}
