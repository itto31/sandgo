package routers

import (
	"net/http"

	"github.com/itto31/sandgo/bd"
)

//EliminarTweet elimina el tweet
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Deve enviar el parametro ID", http.StatusBadRequest)
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar el twwet", http.StatusBadRequest)
	}
	w.Header().Set("content-tyoe", "application/json")
	w.WriteHeader(http.StatusCreated)
}
