package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

func SelectIssueAffect(c *gin.Context) {
	// Params
	option := entity.IssueAffectOption{}
	c.ShouldBind(&option)

	// Action
	issueAffects, err := repository.SelectIssueAffect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": issueAffects})
}

func SelectIssueAffectResult(c *gin.Context) {
	var enumResult = struct {
		AffectResultResultUnKnown entity.AffectResultResult
		AffectResultResultYes     entity.AffectResultResult
		AffectResultResultNo      entity.AffectResultResult
	}{
		AffectResultResultUnKnown: entity.AffectResultResultUnKnown,
		AffectResultResultYes:     entity.AffectResultResultYes,
		AffectResultResultNo:      entity.AffectResultResultNo,
	}

	c.JSON(200, gin.H{"data": enumResult})
}

func CreateOrUpdateIssueAffect(c *gin.Context) {
	// Params
	issueAffect := entity.IssueAffect{}
	c.ShouldBind(&issueAffect)

	// Action
	err := service.CreateOrUpdateIssueAffect(&issueAffect)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
