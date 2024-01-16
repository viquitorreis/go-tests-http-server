package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Buscando o score do jogador Victor", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/victor", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("Recebido %q esperava %q", got, want)
		}
	})
}
