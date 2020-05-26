package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/models"
	"github.com/solrac97gr/yendoapi/utilities"
)

/*CreateWork : Save a user in the database*/
func CreateWork(w http.ResponseWriter, r *http.Request) {

	var work models.Work

	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		http.Error(w, "Bad request "+err.Error(), 400)
		return
	}

	/*Validation*/
	if len(work.Name) == 0 {
		http.Error(w, "Name is required", 400)
		return
	}
	if len(work.Description) == 0 {
		http.Error(w, "Description is required", 400)
		return
	}
	if work.Price <= 20 && work.Price >= 500 {
		http.Error(w, "Price range only in 20 and 500", 400)
		return
	}

	if len(work.AddressID) == 0 {
		http.Error(w, "Address ID is required", 400)
		return
	}
	if len(work.CategoryID) == 0 {
		http.Error(w, "Category ID is required", 400)
		return
	}

	work.CreateAt = time.Now()
	work.UserID = utilities.UserID

	_, isCreated, err := database.RegisterWork(work)
	if err != nil {
		http.Error(w, "Error at moment to register in the db "+err.Error(), 400)
		return
	}
	if !isCreated {
		http.Error(w, "Can't insert the new work", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
