package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateReleaseVersion(c *gin.Context) {
	// Params
	releaseVersion := entity.ReleaseVersion{}
	if err := c.ShouldBindWith(&releaseVersion, binding.JSON); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Action
	err := repository.CreateReleaseVersion(&releaseVersion)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func UpdateReleaseVersion(c *gin.Context) {
	// Params
	releaseVersion := entity.ReleaseVersion{}
	if err := c.ShouldBindWith(&releaseVersion, binding.JSON); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Action
	err := repository.UpdateReleaseVersion(&releaseVersion)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func SelectReleaseVersion(c *gin.Context) {
	// Params
	option := entity.ReleaseVersionOption{}
	if err := c.ShouldBindWith(&option, binding.Form); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Action
	releaseVersions, err := repository.SelectReleaseVersion(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": releaseVersions})
}

func SelectReleaseVersionStatus(c *gin.Context) {
	var enumResult = struct {
		ReleaseVersionStatusOpen     entity.ReleaseVersionStatus
		ReleaseVersionStatusClosed   entity.ReleaseVersionStatus
		ReleaseVersionStatusFrozen   entity.ReleaseVersionStatus
		ReleaseVersionStatusReleased entity.ReleaseVersionStatus
	}{
		ReleaseVersionStatusOpen:     entity.ReleaseVersionStatusOpen,
		ReleaseVersionStatusClosed:   entity.ReleaseVersionStatusClosed,
		ReleaseVersionStatusFrozen:   entity.ReleaseVersionStatusFrozen,
		ReleaseVersionStatusReleased: entity.ReleaseVersionStatusReleased,
	}

	c.JSON(200, gin.H{"data": enumResult})
}

func SelectReleaseVersionType(c *gin.Context) {
	var enumResult = struct {
		ReleaseVersionTypeMajor entity.ReleaseVersionType
		ReleaseVersionTypeMinor entity.ReleaseVersionType
		ReleaseVersionTypePatch entity.ReleaseVersionType
	}{
		ReleaseVersionTypeMajor: entity.ReleaseVersionTypeMajor,
		ReleaseVersionTypeMinor: entity.ReleaseVersionTypeMinor,
		ReleaseVersionTypePatch: entity.ReleaseVersionTypePatch,
	}

	c.JSON(200, gin.H{"data": enumResult})
}
