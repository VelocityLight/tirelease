package controller

import (
	"tirelease/internal/dto"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func SelectIssueRelationInfos(c *gin.Context) {
	// Params
	option := dto.IssueRelationInfoQuery{}
	if err := c.ShouldBindWith(&option, binding.Form); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Action
	issueRelationInfos, err := service.SelectIssueRelationInfo(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": issueRelationInfos})
}
