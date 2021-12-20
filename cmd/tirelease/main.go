package main

import (
	"net/http"
	"tirelease/configs"
	"tirelease/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	configs.LoadConfig("configs/profiles/config.yaml")

	// Connect database
	database.Initialize()

	// Start gin & listen from website
	r := gin.Default()
	r.StaticFS("/", http.Dir("website/build/"))
	r.Run(":8080")
}
