package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
)

/*CreatePay : Create pay endpoint*/
func CreatePay(c *fiber.Ctx){
	c.Send("Create Pay")
	c.SendStatus(http.StatusAccepted)
}