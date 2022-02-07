package controller

import (
	"log"
	"tirelease/internal/entity"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

func UpdateIssueAffectAndTriage(c *gin.Context) {
	option := entity.IssueAffectUpdateOption{}
	if err := c.BindJSON(&option); err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
		return
	}
	if err := service.IssueAffectOperate(&option); err != nil {
		log.Fatal(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
