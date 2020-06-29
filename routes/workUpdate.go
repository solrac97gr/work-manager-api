package routes

import (
	"github.com/gofiber/fiber"
	work2 "github.com/solrac97gr/yendoapi/database/work"
	"net/http"

	"github.com/solrac97gr/yendoapi/models"
)

/*UpdateWork : Edit a user profile*/
func UpdateWork(c *fiber.Ctx) {
	var work models.Work
	var isUpdated bool

	err := c.BodyParser(&work)
	if err != nil {
		c.Send("Bad request" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	isUpdated, err = work2.UpdateWork(work)
	if err != nil {
		c.Send("Error occurred"+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if !isUpdated {
		c.Send("Can't modify the work register")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	c.SendStatus(http.StatusCreated)
}
