package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
)

/*Home: Get default principal route*/
func Home(c *fiber.Ctx) {
	if err := c.JSON(fiber.Map{
		"message":     "Hello Yendo!",
		"description": "Register in the api and then login for use",
		"status":      http.StatusOK,
	}); err != nil {
		c.Status(500).Send(err)
		return
	}
}
