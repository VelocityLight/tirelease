package service

import (
	"testing"

	"tirelease/commons/git"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// Connect
	git.Connect("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")

	// List all repositories for the authenticated user
	triageItems, err := CollectTriageItemByRepo("VelocityLight", "tirelease")

	// Assert
	assert.Equal(t, true, len(triageItems) > 0)
	assert.Equal(t, true, err == nil)
}

func TestLabel(t *testing.T) {
	// Connect
	git.Connect("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")

	labels := []string{"jc_test"}
	err := AddLabelOfAccept("VelocityLight", "tirelease", 1, labels)
	
	// Assert
	assert.Equal(t, true, err == nil)
}
