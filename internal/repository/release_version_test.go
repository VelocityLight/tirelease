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
		Major:       5,
		Minor:       4,
		Patch:       1,
		Description: "Patch版本5.4.1",
		Owner:       "jcye",
		Type:        entity.ReleaseVersionTypePatch,
		Status:      entity.ReleaseVersionStatusPlanned,

		Repos:  &[]string{"pingcap/tidb"},
		Labels: &[]string{"affects-5.4"},
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

INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'5.0.7', 5, 0, 7, '', 'Patch版本5.0.7','尹苏', 'Patch', 'planned', "release-5.0");
INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'5.1.5', 5, 1, 5, '', 'Patch版本5.1.5','尹苏', 'Patch', 'planned', "release-5.1");
INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'5.2.4', 5, 2, 4, '', 'Patch版本5.2.4','尹苏', 'Patch', 'planned', "release-5.2");
INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'5.3.1', 5, 3, 1, '', 'Patch版本5.3.1','尹苏', 'Patch', 'released', "release-5.3");
INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'5.4.1', 5, 4, 1, '', 'Patch版本5.4.1','尹苏', 'Patch', 'planned', "release-5.4");
INSERT INTO release_version (create_time, update_time, name, major, minor, patch, addition, description, owner, type, status, release_branch) VALUES (Now(),Now(),'6.0.0', 6, 0, 0, '', 'Major版本6.0.0','尹苏', 'Major', 'upcoming', "release-6.0");

**/
