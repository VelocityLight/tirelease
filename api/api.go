package api

import (
	"net/http"
	"tirelease/internal/controller"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Create gin-routers
func Routers(file string) (router *gin.Engine) {
	router = gin.Default()

	// Cors
	routeCors(router)

	// Root html & Folder static/ for JS/CSS & home/*any for website('url' match to website/src/routes.js)
	routeHtml(router, file)

	// Real REST-API registry
	routeRestAPI(router)

	return router
}

func routeCors(router *gin.Engine) {
	router.Use(Cors())
}

func routeHtml(router *gin.Engine, file string) {
	router.Use(
		static.Serve("/", static.LocalFile(file, true)),
	)
	router.Use(
		static.Serve("/static", static.LocalFile(file, true)),
	)
	homePages := router.Group("/home")
	{
		homePages.GET("/*any", func(c *gin.Context) {
			c.FileFromFS("/", http.Dir(file))
		})
	}
}

// REST-API Function
func routeRestAPI(router *gin.Engine) {
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

	issue := router.Group("/issue")
	{
		issue.GET("", controller.ListIssueInfo)
		issue.POST("/affect", controller.UpdateIssueAffectAndTriage)

	}

}
