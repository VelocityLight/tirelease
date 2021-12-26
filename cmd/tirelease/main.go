package main

import (
	"tirelease/api"
	"tirelease/commons/configs"
	"tirelease/commons/database"
	"tirelease/commons/git"
)

func main() {
	// Load config
	configs.LoadConfig("config.yaml")

	// Connect database
	database.Connect()

	// Github Client
	git.Connect(configs.Config.Github.AccessToken)

	// Start website & rest api
	router := api.Routers("website/build/")
	router.Run(":8080")
}
