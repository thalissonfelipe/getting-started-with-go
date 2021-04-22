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
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
