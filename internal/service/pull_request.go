package service

import (
	"regexp"
	"strconv"

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

func GetPullRequestRefIssuesByRegexFromV4(pr *git.PullRequestField) ([]int, error) {
	issueNumbers := make([]int, 0)
	if pr == nil {
		return issueNumbers, nil
	}

	// from body reference
	if pr.Body != "" {
		bodyIssues, err := RegexReferenceNumbers(string(pr.Body))
		if err != nil {
			return nil, err
		}
		issueNumbers = append(issueNumbers, bodyIssues...)
	}

	// from timeline reference
	edges := pr.TimelineItems.Edges
	if nil != edges && len(edges) > 0 {
		for i := range edges {
			edge := edges[i]
			if nil == &edge.Node || nil == &edge.Node.IssueComment ||
				nil == &edge.Node.IssueComment.Body {
				continue
			}
			if git.IssueComment != edge.Node.Typename {
				continue
			}
			issueComment := string(edge.Node.IssueComment.Body)
			issueCommentNumbers, err := RegexReferenceNumbers(issueComment)
			if err != nil {
				return nil, err
			}
			issueNumbers = append(issueNumbers, issueCommentNumbers...)
		}
	}

	return issueNumbers, nil
}

func RegexReferenceNumbers(text string) ([]int, error) {
	// param protect
	issueNumbers := make([]int, 0)
	if text == "" {
		return issueNumbers, nil
	}

	// issue number regex
	issueStrs := make([]string, 0)
	re := regexp.MustCompile(`[#][0-9]+`)
	for _, match := range re.FindAllString(text, -1) {
		re2 := regexp.MustCompile(`[0-9]+`)
		issueStrs = append(issueStrs, re2.FindAllString(match, -1)...)
	}

	// compose issue number list
	if len(issueStrs) > 0 {
		for _, issueStr := range issueStrs {
			issueNumber, err := strconv.Atoi(issueStr)
			if nil != err {
				return nil, err
			}
			issueNumbers = append(issueNumbers, issueNumber)
		}
	}

	return issueNumbers, nil
}
