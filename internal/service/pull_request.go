package service

import (
	"tirelease/commons/git"
	"tirelease/internal/entity"
)

// Query PullRequest From Github And Construct Issue Data Service
func GetPullRequestByNumberFromV3(owner, repo string, number int) (*entity.PullRequest, error) {
	pr, _, err := git.Client.GetPullRequestByNumber(owner, repo, number)
	if nil != err {
		return nil, err
	}
	return entity.ComposePullRequestFromV3(pr), nil
}
