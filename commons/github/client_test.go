package github

import (
	"context"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// Connect
	Connect("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")

	// list all repositories for the authenticated user
	repos, _, err := Client.Client.Repositories.List(context.Background(), "", nil)

	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(repos) > 0)
}
