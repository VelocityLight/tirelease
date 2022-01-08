package controller

import (
	"log"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

func ListIssueInfo(c *gin.Context) {
	state := c.DefaultQuery("state", "CLOSED")
	issueInfos, err := service.ListIssueInfo(state)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, issueInfos)
}
