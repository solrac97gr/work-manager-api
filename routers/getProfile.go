package routers

import (
	"encoding/json"
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
)

/*GetProfile : Get the user profile*/
func GetProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	profile, err := database.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error Ocurred"+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&profile)
}
