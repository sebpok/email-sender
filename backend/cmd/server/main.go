package main

import (
	"log"
	"net/http"

	v1 "github.com/sebpok/email-sender/internal/api/v1"
	"github.com/sebpok/email-sender/internal/database"
)

func main() {
	database.Connect()

	r := v1.NewRouter()
	log.Println("Server started on: localhost:8080")
	http.ListenAndServe(":8080", r)
}
