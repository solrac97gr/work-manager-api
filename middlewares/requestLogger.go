package middlewares

import (
	"github.com/gofiber/fiber"
	"log"
)

/*CheckToken : Check the validate of the jwt*/
func RequestLogger(c *fiber.Ctx) {
	method := string(c.Fasthttp.Request.Header.Method())
	url:= string(c.Fasthttp.Request.Header.RequestURI())
	log.Println(method,url)
	c.Next()
}

