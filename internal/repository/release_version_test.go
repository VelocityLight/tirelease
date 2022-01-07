package repository

import (
	"testing"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestReleaseVersion(t *testing.T) {
	// Init
	var config = generateConfig()
	database.Connect(config)

	// Create
	var version = &entity.ReleaseVersion{
		ID: 100,

		CreateTime:        time.Now(),
		UpdateTime:        time.Now(),
		PlanReleaseTime:   time.Now(),
		ActualReleaseTime: time.Now(),

		Name:        "5.4.1",
		Description: "Patch版本5.4.1",
		Owner:       "jcye",
		Type:        entity.ReleaseVersionTypePatch,
		Status:      entity.ReleaseVersionStatusOpen,

		FatherReleaseVersionName: "5.4.0",
		Repos:                    &[]string{"pingcap/tidb"},
		Labels:                   &[]string{"affects-5.4"},
	}
	err := CreateReleaseVersion(version)
	// Assert
	assert.Equal(t, true, err == nil)

	// Update
	version.Status = entity.ReleaseVersionStatusReleased
	err = UpdateReleaseVersion(version)
	// Assert
	assert.Equal(t, true, err == nil)

	// Select
	var option = &entity.ReleaseVersionOption{
		Name: "5.4.1",
	}
	versions, err := SelectReleaseVersion(option)
	// Assert
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(*versions) > 0)
}
