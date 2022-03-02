package git

import "github.com/shurcooL/githubv4"

// ============================================================= Constants
var CrossReferencedEvent = "CrossReferencedEvent"

// ============================================================= Struct Of Needed Fields
type IssueField struct {
	IssueFieldWithoutTimelineItems
	TimelineItems IssueTimelineItems `graphql:"timelineItems(first: 50, itemTypes: [CROSS_REFERENCED_EVENT, CLOSED_EVENT] )"`
}

type PullRequestField struct {
	PullRequestFieldWithoutTimelineItems
	TimelineItems PullRequestTimelineItems `graphql:"timelineItems(first: 50, itemTypes: [CROSS_REFERENCED_EVENT, ISSUE_COMMENT] )"`
}

type RepositoryField struct {
	Name  githubv4.String
	Owner struct {
		Login githubv4.String
	}
}

type LabelField struct {
	Nodes []struct {
		Name githubv4.String
	}
}

type AssigneesFiled struct {
	Nodes []struct {
		Login     githubv4.String
		CreatedAt githubv4.DateTime
	}
}

type UserField struct {
	Login githubv4.String
}

type IssueFieldWithoutTimelineItems struct {
	Title      githubv4.String
	State      githubv4.IssueState
	ID         githubv4.ID
	Number     githubv4.Int
	Url        githubv4.String
	Author     UserField
	Body       githubv4.String
	ClosedAt   githubv4.DateTime
	CreatedAt  githubv4.DateTime
	UpdatedAt  githubv4.DateTime
	Repository RepositoryField
	Labels     LabelField     `graphql:"labels(first: 30)"`
	Assignees  AssigneesFiled `graphql:"assignees(first: 10)"`
}

type PullRequestFieldWithoutTimelineItems struct {
	ID          githubv4.ID
	Number      githubv4.Int
	State       githubv4.PullRequestState
	Title       githubv4.String
	Repository  RepositoryField
	Url         githubv4.String
	HeadRefName githubv4.String

	CreatedAt githubv4.DateTime
	UpdatedAt githubv4.DateTime
	ClosedAt  *githubv4.DateTime
	MergedAt  *githubv4.DateTime

	Merged    githubv4.Boolean
	Mergeable githubv4.MergeableState

	MergeCommit struct {
		OID           githubv4.GitObjectID
		CommittedDate githubv4.DateTime
	}
	Author UserField

	Labels    LabelField     `graphql:"labels(first: 30)"`
	Assignees AssigneesFiled `graphql:"assignees(first: 10)"`

	BaseRefName githubv4.String
}

type IssueTimelineItems struct {
	Edges []struct {
		Node struct {
			Typename             string `graphql:"__typename"`
			CrossReferencedEvent struct {
				WillCloseTarget githubv4.Boolean
				Source          struct {
					PullRequest PullRequestField `graphql:"... on PullRequest"`
				}
			} `graphql:"... on CrossReferencedEvent"`
			ClosedEvent struct {
				Actor  UserField
				Closer struct {
					PullRequest PullRequestField `graphql:"... on PullRequest"`
				}
			} `graphql:"... on ClosedEvent"`
		}
	}
}

type PullRequestTimelineItems struct {
	Edges []struct {
		Node struct {
			Typename             string `graphql:"__typename"`
			CrossReferencedEvent struct {
				Source struct {
					PullRequest PullRequestFieldWithoutTimelineItems `graphql:"... on PullRequest"`
				}
			} `graphql:"... on CrossReferencedEvent"`
			IssueComment struct {
				Author UserField
				Body   githubv4.String
			} `graphql:"... on IssueComment"`
		}
	}
}
