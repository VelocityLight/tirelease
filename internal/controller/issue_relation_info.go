package controller

import (
	"fmt"

	"tirelease/internal/dto"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

func SelectIssueRelationInfos(c *gin.Context) {
	// Params
	option := dto.IssueRelationInfoQuery{}
	if err := c.ShouldBind(&option); err != nil {
		c.JSON(500, err.Error())
		return
	}
	fmt.Println(option)

	// Action
	issueRelationInfos, err := service.SelectIssueRelationInfo(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": issueRelationInfos})
}
