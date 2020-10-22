package main

import (
	"log"

	"twittor/bd"
	"twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BDD")
		return
	}
	handlers.Manejadores()
}
