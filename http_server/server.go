package http_server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStorer interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStorer
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}
