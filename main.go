package main

import (
	"log"

	"github.com/solrac97gr/twitter-fake/database"
	"github.com/solrac97gr/twitter-fake/handlers"
)

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not conected to DB")
		return
	}
	handlers.Handlers()
}
