package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	middlew "github.com/itto31/sandgo/middleW"
	"github.com/itto31/sandgo/routers"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el Handler y pongo a escuchar al servicio
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	//router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJwt(routers.Verperfil)).Methods("POST"))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
