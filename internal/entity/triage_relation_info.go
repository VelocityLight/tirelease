package entity

// TriageRelationInfo Struct
type TriageRelationInfo struct {
	Issue            *Issue
	IssueAffects     []*IssueAffect
	IssuePrRelations []*IssuePrRelation
	PullRequests     []*PullRequest
}
