package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Rest api registry & Create gin-routers
func Routers(file string) (router *gin.Engine) {
	router = gin.New()

	router.Use(
		static.Serve("/", static.LocalFile(file, true)),
	)

	v1 := router.Group("/ping")
	{
		v1.GET("/", pong)
	}

	return router
}
