package middlewares

import (
	"github.com/gofiber/fiber"
	"net/http"

	"github.com/solrac97gr/yendoapi/utilities"
)

/*CheckToken : Check the validate of the jwt*/
func CheckToken(c *fiber.Ctx) {
	_, _, _, err := utilities.ProcessToken(c.Get("Authorization"))
	if err != nil {
		c.Send("Error in the jwt !"+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	c.Next()
}
