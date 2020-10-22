package main

import (
	"log"

	"github.com/asaavedras-dev/twittor/bd"
	"github.com/asaavedras-dev/twittor/handlers"
)

func main() {
	if bd.CheaqueoConnection() == 0 {
		log.Fatal("Sin conexión a la BDD")
		return
	}
	handlers.Manejadores()
}
