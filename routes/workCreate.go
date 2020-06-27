package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
	"time"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/models"
	"github.com/solrac97gr/yendoapi/utilities"
)

/*CreateWork : Save a user in the database*/
func CreateWork(c *fiber.Ctx) {

	var work models.Work

	err := c.BodyParser(work)
	if err != nil {
		c.Send("Bad request " + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	/*Validation*/
	if len(work.Name) == 0 {
		c.Send("Name is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if len(work.Description) == 0 {
		c.Send("Description is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if work.Price <= 20 && work.Price >= 500 {
		c.Send("Price range only in 20 and 500")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	if len(work.AddressID) == 0 {
		c.Send("Address ID is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if len(work.CategoryID) == 0 {
		c.Send("Category ID is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	work.CreateAt = time.Now()
	work.UserID = utilities.UserID

	_, isCreated, err := database.RegisterWork(work)
	if err != nil {
		c.Send("Error at moment to register in the db "+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if !isCreated {
		c.Send("Can't insert the new work")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	c.SendStatus(http.StatusCreated)
}
