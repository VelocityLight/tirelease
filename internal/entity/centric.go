package entity

// Data-Centric return struct
type IssueCentric struct {
	Issue        *Issue
	PullRequests []*PullRequest
}

type PullRequestCentric struct {
	PullRequests *PullRequest
	Issue        []*Issue
}
