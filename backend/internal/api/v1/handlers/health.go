package handlers

import "net/http"

func CheckHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server status: ok"))
}
