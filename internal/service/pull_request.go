package service

import (
	"tirelease/commons/git"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
)

// GetPullRequestFromV3
func GetPullRequestByNumberFromV3(owner, repo string, number int) (*entity.PullRequest, error) {
	pr, _, err := git.Client.GetPullRequestByNumber(owner, repo, number)
	if nil != err {
		return nil, err
	}
	return ConsistPullRequestFromV3(pr), nil
}

// ConsistPullRequestFromV3
func ConsistPullRequestFromV3(pullRequest *github.PullRequest) *entity.PullRequest {
	labels := &[]github.Label{}
	for _, node := range pullRequest.Labels {
		*labels = append(*labels, *node)
	}
	assignees := &[]github.User{}
	for _, node := range pullRequest.Assignees {
		*assignees = append(*assignees, *node)
	}
	requestedReviewers := &[]github.User{}
	for _, node := range pullRequest.RequestedReviewers {
		*requestedReviewers = append(*requestedReviewers, *node)
	}

	return &entity.PullRequest{
		PullRequestID: *pullRequest.NodeID,
		Number:        *pullRequest.Number,
		State:         *pullRequest.State,
		Title:         *pullRequest.Title,
		Owner:         *pullRequest.Base.Repo.Owner.Login,
		Repo:          *pullRequest.Base.Repo.Name,
		HTMLURL:       *pullRequest.HTMLURL,
		HeadBranch:    *pullRequest.Head.Ref,

		CreatedAt: *pullRequest.CreatedAt,
		UpdatedAt: *pullRequest.UpdatedAt,
		ClosedAt:  pullRequest.ClosedAt,
		MergedAt:  pullRequest.MergedAt,

		Merged:         *pullRequest.Merged,
		Mergeable:      pullRequest.Mergeable,
		MergeableState: pullRequest.MergeableState,

		Labels:             labels,
		Assignee:           pullRequest.Assignee,
		Assignees:          assignees,
		RequestedReviewers: requestedReviewers,
	}
}
