package main

import (
	"adpc-webserver/src/database"
	"adpc-webserver/src/routes"
	"log"
)

func main() {
	_, err := database.ConectaDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	routes.HandleRequest()
}
