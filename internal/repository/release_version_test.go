package repository

import (
	"testing"

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
		Name:        "5.4.1",
		Description: "Patch版本5.4.1",
		Owner:       "jcye",
		Type:        entity.ReleaseVersionTypePatch,
		Status:      entity.ReleaseVersionStatusOpen,

		FatherReleaseVersionName: "5.4",
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

/**
sql:

INSERT INTO release_version (create_time, update_time, name, description, owner, type, status, father_release_version_name, release_branch) VALUES (Now(),Now(),'5.4','Minor版本5.4.0','JunChenYe', 'Minor', 'Open', '5.4', "release-5.4");
INSERT INTO release_version (create_time, update_time, name, description, owner, type, status, father_release_version_name, release_branch) VALUES (Now(),Now(),'5.4.1','Patch版本5.4.1','JunChenYe', 'Patch', 'Open', '5.4', "release-5.4");

**/
