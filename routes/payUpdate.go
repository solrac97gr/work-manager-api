package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
)

/*UpdatePay: Update pay*/
func UpdatePay(c *fiber.Ctx){
	c.Send("Update Pay")
	c.SendStatus(http.StatusAccepted)
}
