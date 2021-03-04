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
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJwt(routers.Verperfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJwt(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJwt(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJwt(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJwt(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJwt(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJwt(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJwt(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJwt(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJwt(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJwt(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJwt(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
