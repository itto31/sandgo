package middlew

import (
	"net/http"

	"github.com/itto31/sandgo/bd"
)

//ChequeoBD revisa la conexion con la base
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion pedida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
