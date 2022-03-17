package controller

import (
	"log"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Rest-API controller
func WebhookHandler(c *gin.Context) {
	webhookPayload := service.WebhookPayload{}
	if err := c.ShouldBindWith(&webhookPayload, binding.JSON); err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
		return
	}
	if err := service.UpdatePrAndIssue(webhookPayload); err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func InitDataForDemo(c *gin.Context) {
	service.InitDB()
}
