package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdateVersionTriage(c *gin.Context) {
	// Params
	versionTriage := &entity.VersionTriage{}
	c.ShouldBind(versionTriage)

	// Action
	versionTriageInfo, err := service.CreateOrUpdateVersionTriageInfo(versionTriage)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	// Response
	var statusCode int = 200
	if nil != versionTriageInfo && versionTriageInfo.IsFrozen && versionTriageInfo.IsAccept {
		statusCode = 202
	}
	c.JSON(statusCode, gin.H{"status": "ok", "data": versionTriageInfo})
}
