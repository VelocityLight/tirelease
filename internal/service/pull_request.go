package service

import (
	"tirelease/commons/git"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

// Operation
func AddLabelByPullRequestID(pullRequestID, label string) error {
	// select issue by id
	option := &entity.PullRequestOption{
		PullRequestID: pullRequestID,
	}
	pr, err := repository.SelectPullRequestUnique(option)
	if nil != err {
		return err
	}

	// add issue label
	_, _, err = git.Client.AddLabel(pr.Owner, pr.Repo, pr.Number, label)
	if nil != err {
		return err
	}
	return nil
}

func RemoveLabelByPullRequestID(pullRequestID, label string) error {
	// select issue by id
	option := &entity.PullRequestOption{
		PullRequestID: pullRequestID,
	}
	pr, err := repository.SelectPullRequestUnique(option)
	if nil != err {
		return err
	}

	// remove issue label
	_, err = git.Client.RemoveLabel(pr.Owner, pr.Repo, pr.Number, label)
	if nil != err {
		return err
	}
	return nil
}

// Query PullRequest From Github And Construct Issue Data Service
func GetPullRequestByNumberFromV3(owner, repo string, number int) (*entity.PullRequest, error) {
	pr, _, err := git.Client.GetPullRequestByNumber(owner, repo, number)
	if nil != err {
		return nil, err
	}
	return entity.ComposePullRequestFromV3(pr), nil
}
