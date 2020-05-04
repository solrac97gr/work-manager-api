package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/solrac97gr/yendoapi/middlewares"
	"github.com/solrac97gr/yendoapi/routers"
)

/*Handlers : set the port,cors,handlers and then serve the api*/
func Handlers() {
	router := mux.NewRouter()
	/*User : Login and Register routes*/
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	/*Profile : Get and Create*/
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.CheckToken(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.CheckToken(routers.UpdateProfile))).Methods("PUT")
	/*Work :Complete CRUD */
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routers.CreateWork))).Methods("POST")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routers.UpdateWork))).Methods("PUT")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routers.CreateWork))).Methods("DELETE")
	router.HandleFunc("/work", middlewares.CheckDB(middlewares.CheckToken(routers.CreateWork))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Server run in http://localhost:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
