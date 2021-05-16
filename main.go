package main

import (
	"github.com/TeamKun/mcapi/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handler.RootHandler).Methods("GET")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
