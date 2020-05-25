package middlewares

import (
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
)

/*CheckDB : Check the DB connection before to execute a handle func*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !database.ConnectionOK() {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
