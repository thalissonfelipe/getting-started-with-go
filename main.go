package main

import (
	"log"
	"net/http"

	hs "github.com/thalissonfelipe/learn-go/http_server"
)

func main() {
	server := &hs.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
