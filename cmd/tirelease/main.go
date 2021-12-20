package main

import (
	"tirelease/api"
	"tirelease/configs"
	"tirelease/internal/database"
)

func main() {
	// Load config
	configs.LoadConfig("configs/profiles/config.yaml")

	// Connect database
	database.Connect()

	// Start website & rest api
	router := api.Routers("website/build/")
	router.Run(":8080")
}
