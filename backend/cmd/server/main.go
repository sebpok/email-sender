package main

import (
	"net/http"

	v1 "github.com/sebpok/email-sender/internal/api/v1"
)

func main() {
	r := v1.NewRouter()
	http.ListenAndServe(":8080", r)
}
