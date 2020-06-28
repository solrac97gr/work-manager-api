package routes

import "github.com/gofiber/fiber"

func GetWorks(c* fiber.Ctx){
	c.Send("Get one o")
}
