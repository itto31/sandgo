package main

import (
	"log"

	"github.com/itto31/sandgo/bd"
	"github.com/itto31/sandgo/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
