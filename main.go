package main

import (
	"github.com/solrac97gr/yendoapi/database"
	"github.com/solrac97gr/yendoapi/handlers"
	"log"
)

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not connected to DB")
		return
	}
	handlers.Handlers()
}
