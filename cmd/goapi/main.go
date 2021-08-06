package main

import (
	"log"
	"net/http"

	"github.com/trastolillo/go-rest-api/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
