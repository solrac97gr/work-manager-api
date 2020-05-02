package routers

import (
	"encoding/json"
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/models"
)

/*UpdateProfile : Edit a user profile*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var isUpdated bool

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request"+err.Error(), 400)
		return
	}

	isUpdated, err = database.ModifyProfile(user, UserID)
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
