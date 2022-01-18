package entity

// Data-Center return struct
type IssueDataCenter struct {
	Issue        *Issue
	PullRequests []*PullRequest
}

type PullRequestDataCenter struct {
	PullRequests *PullRequest
	Issue        []*Issue
}
