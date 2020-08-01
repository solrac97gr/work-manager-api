package main

import (
	"github.com/solrac97gr/yendoapi/api"
	"github.com/solrac97gr/yendoapi/database"
	"log"
)

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not connected to DB")
		return
	}
	api.Init()
}
