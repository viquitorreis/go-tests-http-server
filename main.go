package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server)) // se tiver algum erro ao servir nosso servidor, o '1og' vai reportar o erro e vai terminar o programa
}
