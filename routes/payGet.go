package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
)

/*GetPay: Get all pays*/
func GetPay(c *fiber.Ctx){
	c.Send("Get Pay")
	c.SendStatus(http.StatusAccepted)
}