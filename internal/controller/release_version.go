package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
)

func CreateReleaseVersion(c *gin.Context) {
	// Params
	releaseVersion := &entity.ReleaseVersion{}
	c.ShouldBind(releaseVersion)

	// Action
	err := repository.CreateReleaseVersion(releaseVersion)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func UpdateReleaseVersion(c *gin.Context) {
	// Params
	releaseVersion := &entity.ReleaseVersion{}
	c.ShouldBind(releaseVersion)

	// Action
	err := repository.UpdateReleaseVersion(releaseVersion)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func SelectReleaseVersion(c *gin.Context) {
	// Params
	option := entity.ReleaseVersionOption{}
	c.ShouldBind(&option)

	// Action
	releaseVersions, err := repository.SelectReleaseVersion(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": releaseVersions})
}
