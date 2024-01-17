package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Println(player)

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	// r.URL.Path retorna o path da request na qual podemos depois usar o strings.TrimPrefix para “apararmos / removermos” o /players/ e pegar o player desejado
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		fmt.Println("Nenhum score")
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Println(score)
	// ATENÇÃO !!! Fprint != Fprintf
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
