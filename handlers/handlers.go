package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/solrac97gr/yendoapi/middlewares"
	"github.com/solrac97gr/yendoapi/routes"
)

/*Handlers : set the port,cors,handlers and then serve the api*/
func Handlers() {
	router := mux.NewRouter()
	/*User : Login and Register routes*/
	router.HandleFunc("/register", middlewares.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routes.Login)).Methods("POST")
	/*Profile : Get and Edit*/
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.CheckToken(routes.GetProfile))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.CheckToken(routes.UpdateProfile))).Methods("PUT")
	/*Work :Complete CRUD */
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routes.CreateWork))).Methods("POST")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routes.UpdateWork))).Methods("PUT")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routes.CreateWork))).Methods("DELETE")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routes.CreateWork))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Server run in http://localhost:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
