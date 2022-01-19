package main

import (
	"tirelease/api"
	"tirelease/commons/configs"
	"tirelease/commons/database"
	"tirelease/commons/git"
	"tirelease/internal/cron"
)

func main() {
	// Load config
	configs.LoadConfig("config.yaml")

	// Connect database
	database.Connect(configs.Config)

	// Github Client (If Needed)
	git.ConnectV4(configs.Config.Github.AccessToken)

	// Start Cron (If Needed)
	cron.DemoCron()

	// Start website && REST-API
	router := api.Routers("website/build/")
	router.Run(":8080")
}
