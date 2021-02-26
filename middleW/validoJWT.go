package middlew

import (
	"net/http"

	"github.com/itto31/sandgo/routers"
)

//ValidoJwt permite validar el JWT que nos viene en la peticion
func ValidoJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error en el Token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
