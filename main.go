package main

import (
	"log"

	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/handlers"
)

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not connected to DB")
		return
	}
	handlers.Handlers()
}
