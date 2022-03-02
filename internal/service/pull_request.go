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

// ConsistPullRequestFromV4
// TODO: v4 implement by tony at 2022/02/14
func ConsistPullRequestFromV4(pullRequestField *git.PullRequestField) *entity.PullRequest {
	labels := &[]github.Label{}
	for _, labelNode := range pullRequestField.Labels.Nodes {
		label := github.Label{
			Name: github.String(string(labelNode.Name)),
		}
		*labels = append(*labels, label)
	}
	assignees := &[]github.User{}
	for _, userNode := range pullRequestField.Assignees.Nodes {
		user := github.User{
			Login:     (*string)(&userNode.Login),
			CreatedAt: (*github.Timestamp)(&userNode.CreatedAt),
		}
		*assignees = append(*assignees, user)
	}

	return &entity.PullRequest{
		PullRequestID: pullRequestField.ID.(string),
		Number:        int(pullRequestField.Number),
		State:         string(pullRequestField.State),
		Title:         string(pullRequestField.Title),
		Owner:         string(pullRequestField.Repository.Owner.Login),
		Repo:          string(pullRequestField.Repository.Name),
		HTMLURL:       string(pullRequestField.Url),
		HeadBranch:    string(pullRequestField.BaseRefName),
		CreatedAt:     pullRequestField.CreatedAt.Time,
		UpdatedAt:     pullRequestField.UpdatedAt.Time,
		Merged:        bool(pullRequestField.Merged),
		Labels:        labels,
		Assignees:     assignees,
	}
}
