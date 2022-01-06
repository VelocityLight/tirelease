package git

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubInfoV4 struct {
	Client *githubv4.Client
}

// V4版本GraphAPI
var ClientV4 = &GithubInfoV4{}

func ConnectV4(accessToken string) {
	// Outh
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	httpClient := oauth2.NewClient(context.Background(), src)

	// Client
	ClientV4.Client = githubv4.NewClient(httpClient)
}
