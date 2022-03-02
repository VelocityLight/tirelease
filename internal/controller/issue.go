package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
)

func SelectIssue(c *gin.Context) {
	// Params
	option := entity.IssueOption{}
	c.ShouldBind(&option)

	// Action
	issues, err := repository.SelectIssue(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": issues})
}
