package middlewares

import (
	"github.com/gofiber/fiber"
	"net/http"

	"github.com/solrac97gr/yendoapi/database"
)

/*CheckDB : Check the DB connection before to execute a handle func*/
func CheckDB(c *fiber.Ctx) {
	if !database.ConnectionOK() {
		c.Send("Conexión perdida con la base de datos")
		c.SendStatus(http.StatusInternalServerError)
		return
	}
	c.Next()
}
