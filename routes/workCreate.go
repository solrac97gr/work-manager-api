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

	err := c.BodyParser(&work)
	if err != nil {
		c.Send("Bad request " + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	work.CreateAt = time.Now()
	work.UserID = utilities.UserID

	/*Validation*/

	err=work.Validate()
	if err != nil {
		c.Send(err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

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
