package routes

import (
	"github.com/gofiber/fiber"
	"net/http"
)

/*Delete: Update pay*/
func DeletePay(c *fiber.Ctx){
	c.Send("Update Pay")
	c.SendStatus(http.StatusAccepted)
}
