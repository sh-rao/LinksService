package main

import (
	"net/http"

	"github.com/gorilla/mux"

	h "./internal"
	l "./internal/link"
)

func main() {
	linkService := l.NewService()
	handler := h.NewHandler(*linkService)
	r := mux.NewRouter()
	r.HandleFunc("/v1/users/{id}", handler.PostHandler).Methods("POST")
	r.HandleFunc("/v1/users/{id}", handler.GetHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
