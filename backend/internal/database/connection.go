package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func Connect() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"user",
		"password",
		"localhost",
		"5432",
		"email_sender",
		"disable",
	)

	c, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB CONN IS WORKING")
	conn = c
}

func GetConnection() *sql.DB {
	return conn
}
