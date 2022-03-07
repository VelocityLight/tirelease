package controller

import (
	"tirelease/internal/dto"
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

func SelectVersionTriageInfos(c *gin.Context) {
	// Params
	versionTriageInfoQuery := &dto.VersionTriageInfoQuery{}
	c.ShouldBind(versionTriageInfoQuery)

	// Action
	versionTriageInfos, err := service.SelectVersionTriageInfo(versionTriageInfoQuery)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	// Response
	c.JSON(200, gin.H{"data": versionTriageInfos})
}
