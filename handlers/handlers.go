package handlers

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/solrac97gr/yendoapi/middlewares"
	"github.com/solrac97gr/yendoapi/routes"
	"log"
)

/*Handlers : set the port,cors,handlers and then serve the api*/
func Handlers() {
	app := fiber.New()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTION"},
		AllowCredentials: true,
	}
	app.Use(cors.New(config))
	app.Use(middlewares.CheckDB)
	app.Get("/", func(c *fiber.Ctx) {
		if err := c.JSON(fiber.Map{
			"message": "Hello Yendo!",
			"status": "online",
		}); err != nil {
			c.Status(500).Send(err)
			return
		}
	})
	/*User : Login and Register routes*/
	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	/*Protected Routes*/
	app.Use(middlewares.CheckToken)
	/*Profile : Get and Edit*/
	app.Get("/profile", routes.GetProfile)
	app.Put("/profile", routes.UpdateProfile)
	/*Work :Complete CRUD */
	app.Post("/work", routes.CreateWork)
	app.Put("/work", routes.UpdateWork)
	app.Delete("/work", routes.CreateWork)
	app.Get("/work", routes.CreateWork)

	err := app.Listen(8080)
	if err != nil {
		log.Fatal(err.Error())
	}
}
