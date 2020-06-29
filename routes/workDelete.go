package routes

import (
	"github.com/gofiber/fiber"
	"github.com/solrac97gr/yendoapi/database"
	"net/http"
)

func DeleteWork(c *fiber.Ctx) {
	ID := c.Query("id")

	if len(ID) < 1 {
		c.Send(ID)
		c.Send("Parameter ID is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	_, err := database.DeleteWork(ID)
	if err != nil {
		c.Send("Error Occurred " + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	c.Accepts("application/json")
	c.Send("Register deleted")
	c.SendStatus(http.StatusAccepted)
}
