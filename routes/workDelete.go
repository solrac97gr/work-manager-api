package routes

import "github.com/gofiber/fiber"

func DeleteWork(c *fiber.Ctx){
	c.Send("Received")
}
