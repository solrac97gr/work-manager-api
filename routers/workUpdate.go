package routers

import (
	"encoding/json"
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/models"
)

/*UpdateWork : Edit a user profile*/
func UpdateWork(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var work models.Work
	var isUpdated bool

	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		http.Error(w, "Bad request"+err.Error(), 400)
		return
	}

	isUpdated, err = database.UpdateWork(work, ID)
	if err != nil {
		http.Error(w, "Error ocurred"+err.Error(), 400)
		return
	}
	if !isUpdated {
		http.Error(w, "Can't modify the user register", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
