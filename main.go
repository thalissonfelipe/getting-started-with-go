package main

import (
	"log"
	"net/http"

	hs "github.com/thalissonfelipe/learn-go/http_server"
)

func main() {
	handler := http.HandlerFunc(hs.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
