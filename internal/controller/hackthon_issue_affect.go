package controller

import (
	"log"
	"tirelease/internal/entity"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UpdateIssueAffectAndTriage(c *gin.Context) {
	option := entity.IssueAffectUpdateOption{}
	if err := c.ShouldBindWith(&option, binding.JSON); err != nil {
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
