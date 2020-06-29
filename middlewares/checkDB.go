package middlewares

import (
	"github.com/gofiber/fiber"
	"github.com/solrac97gr/yendoapi/database"
	"net/http"
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
