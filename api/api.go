package api

import (
	"tirelease/internal/controller"

	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Create gin-routers
func Routers(file string) (router *gin.Engine) {
	router = gin.Default()

	// Cors
	router.Use(Cors())

	// Static fronted file
	// router.Use(
	// 	static.Serve("/page", static.LocalFile(file, true)),
	// )
	// router.LoadHTMLGlob(file + ".html")

	// Test "ping"
	ping := router.Group("/ping")
	{
		ping.GET("/", pong)
	}

	// REST API registry
	testEntity := router.Group("/testentity")
	{
		testEntity.GET("/select", controller.TestEntitySelect)
		testEntity.POST("/insert", controller.TestEntityInsert)
	}
	triageItem := router.Group("/triage")
	{
		triageItem.GET("/select", controller.SelectTriageItems)
		triageItem.POST("/insert", controller.InsertTriageItems)
		triageItem.POST("/accept", controller.AddLabelsToIssue)
	}

	return router
}
