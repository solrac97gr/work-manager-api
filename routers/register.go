package routers

import (
	"encoding/json"
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/models"
)

/*Register : Save a user in the database*/
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request "+err.Error(), 400)
		return
	}
	/*Validation*/
	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password need at less 6 char", 400)
		return
	}
	_, found, _ := database.UserExist(user.Email)
	if found {
		http.Error(w, "Email alredy register", 400)
		return
	}
	_, isCreated, err := database.UserRegister(user)
	if err != nil {
		http.Error(w, "Error at moment to register in the db "+err.Error(), 400)
		return
	}
	if !isCreated {
		http.Error(w, "Can't insert the new user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
