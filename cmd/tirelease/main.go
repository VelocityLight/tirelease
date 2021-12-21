package main

import (
	"tirelease/api"
	"tirelease/commons/configs"
	"tirelease/commons/database"
)

func main() {
	// Load config
	configs.LoadConfig("config.yaml")

	// Connect database
	database.Connect()

	// Start website & rest api
	router := api.Routers("website/build/")
	router.Run(":8080")
}
