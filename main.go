package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
