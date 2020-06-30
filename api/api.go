package api

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/solrac97gr/yendoapi/middlewares"
	"github.com/solrac97gr/yendoapi/routes"
	"log"
)

/*Api : set the port,cors,api and then serve the api*/
func Api() {
	app := fiber.New()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
	}
	app.Use(cors.New(config))
	app.Use(middlewares.CheckDB)
	app.Use(middlewares.RequestLogger)
	app.Get("/", routes.Home)
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
	app.Delete("/work", routes.DeleteWork)
	app.Get("/work", routes.GetWorks)

	err := app.Listen(5000)
	if err != nil {
		log.Fatal(err.Error())
	}
}
