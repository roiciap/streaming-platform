package main

import (
	"net/http"

	"github.com/roiciap/streaming-platform/go/be/user-service/internal/http/handlers"
)

func main() {
	mux := http.NewServeMux()
	handler := handlers.NewUserHandler()
	mux.Handle("/login", handler)
	mux.Handle("/register", handler)

	http.ListenAndServe(":8080", mux)
}
