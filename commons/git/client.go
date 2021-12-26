package git

import (
	"context"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type GithubInfo struct {
	Client *github.Client
}

var Client = &GithubInfo{}

func Connect(accessToken string) {
	// Outh
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Github client
	githubClient := github.NewClient(tc)
	Client.Client = githubClient
}
