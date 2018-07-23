package main

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var connection = DatabaseConnection()

func WriteDomain(domain string, wg *sync.WaitGroup) {
	_, err := connection.Exec(
		"INSERT INTO sources (domain) VALUES ($1)",
		domain)

	if err != nil {
		panic(err)
	}

	wg.Done()
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
