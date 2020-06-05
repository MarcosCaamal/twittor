package main

import (
	"log"

	"github.com/MarcosCaamal/twittor/bd"
	"github.com/MarcosCaamal/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
