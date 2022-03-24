package service

import (
	"testing"
	"tirelease/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestComposeVersionAtom(t *testing.T) {
	str := "5.4.1"
	major, minor, patch, _ := ComposeVersionAtom(str)
	short := ComposeVersionShortType(str)
	assert.Equal(t, 5, major)
	assert.Equal(t, 4, minor)
	assert.Equal(t, 1, patch)
	assert.Equal(t, entity.ReleaseVersionShortTypePatch, short)

	str = "5.4"
	major, minor, patch, _ = ComposeVersionAtom(str)
	short = ComposeVersionShortType(str)
	assert.Equal(t, 5, major)
	assert.Equal(t, 4, minor)
	assert.Equal(t, 0, patch)
	assert.Equal(t, entity.ReleaseVersionShortTypeMinor, short)

	str = "5.4-hotfix-1"
	major, minor, patch, addition := ComposeVersionAtom(str)
	short = ComposeVersionShortType(str)
	assert.Equal(t, 5, major)
	assert.Equal(t, 4, minor)
	assert.Equal(t, 0, patch)
	assert.Equal(t, "hotfix-tiflash-patch1", addition)
	assert.Equal(t, entity.ReleaseVersionShortTypeMinor, short)

	str = "5.4.1-hotfix"
	major, minor, patch, addition = ComposeVersionAtom(str)
	short = ComposeVersionShortType(str)
	assert.Equal(t, 5, major)
	assert.Equal(t, 4, minor)
	assert.Equal(t, 1, patch)
	assert.Equal(t, "hotfix", addition)
	assert.Equal(t, entity.ReleaseVersionShortTypeMinor, short)
}
