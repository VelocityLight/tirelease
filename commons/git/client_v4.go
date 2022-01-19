// Tool Url: https://github.com/shurcooL/githubv4
// Tool Guide: https://docs.github.com/en/graphql

package git

import (
	"context"
	"log"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubInfoV4 struct {
	client *githubv4.Client
}

// V4版本GraphAPI
var ClientV4 = &GithubInfoV4{}

func ConnectV4(accessToken string) {
	// Outh
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	httpClient := oauth2.NewClient(context.Background(), src)

	// Client
	ClientV4.client = githubv4.NewClient(httpClient)
}

func (client *GithubInfoV4) GetIssuesByTimeRange(owner, name string, labels []string, from time.Time, to time.Time, batchLimit int, totalLimit int) (issues []IssueNode, err error) {
	var query struct {
		Repository struct {
			Issues struct {
				Edges []struct {
					Cursor githubv4.String
					Node   IssueNode
				}
			} `graphql:"issues(first: $limit, after: $cursor, orderBy: {field: UPDATED_AT, direction: ASC}, labels: $labels, filterBy: {since: $since})"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	cursor := (*githubv4.String)(nil)
	total := 0
	ghLabels := make([]githubv4.String, 0, len(labels))
	for _, l := range labels {
		ghLabels = append(ghLabels, githubv4.String(l))
	}

	since := from.Add(-1 * time.Minute).Format(time.RFC3339)
	log.Printf("fetching since %s", since)

	for totalLimit != 0 {
		limit := batchLimit
		if totalLimit > 0 && totalLimit < limit {
			limit = totalLimit
		}
		param := map[string]interface{}{
			"name":   githubv4.String(name),
			"owner":  githubv4.String(owner),
			"limit":  githubv4.Int(limit),
			"cursor": cursor,
			"labels": ghLabels,
			"since":  githubv4.String(since),
		}

		err = client.client.Query(context.Background(), &query, param)
		if err != nil {
			log.Println(err)
			return
		}
		edges := query.Repository.Issues.Edges

		for _, edge := range edges {
			issues = append(issues, edge.Node)
			log.Printf("%06d %s %s\n", edge.Node.Number, edge.Node.UpdatedAt.Format(time.RFC3339), edge.Node.Title)
		}

		cnt := len(edges)
		if cnt != 0 {
			lastIssue := &edges[cnt-1]
			cursor = &lastIssue.Cursor
			lastUpdated := lastIssue.Node.UpdatedAt.Time
			total += cnt
			totalLimit -= cnt
			log.Println(cnt, "fetced", owner, name, labels, lastUpdated)
			if lastUpdated.After(to) {
				break
			}
		}
		if cnt != limit {
			break
		}
	}

	log.Printf("fetched %d issues from %s/%s\n", total, owner, name)
	return
}

func (client *GithubInfoV4) GetPullRequestsFrom(owner, name string, from time.Time, batchLimit int, totalLimit int) (prs []PullRequest, err error) {
	var query struct {
		Repository struct {
			PullRequests struct {
				Edges []struct {
					Cursor githubv4.String
					Node   PullRequest
				}
			} `graphql:"pullRequests(first: $limit, after: $cursor, orderBy: {field: UPDATED_AT, direction: DESC})"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	cursor := (*githubv4.String)(nil)
	total := 0

	since := from.Add(-1 * time.Minute)
	log.Printf("fetching since %s", since)

	for totalLimit != 0 {
		limit := batchLimit
		if totalLimit > 0 && totalLimit < limit {
			limit = totalLimit
		}
		param := map[string]interface{}{
			"name":   githubv4.String(name),
			"owner":  githubv4.String(owner),
			"limit":  githubv4.Int(limit),
			"cursor": cursor,
		}

		err = client.client.Query(context.Background(), &query, param)
		if err != nil {
			log.Println(err)
			return
		}
		edges := query.Repository.PullRequests.Edges

		for _, edge := range edges {
			prs = append(prs, edge.Node)
			log.Printf("%06d %s %s\n", edge.Node.Number, edge.Node.UpdatedAt.Format(time.RFC3339), edge.Node.Title)
		}

		cnt := len(edges)
		if cnt != 0 {
			lastIssue := &edges[cnt-1]
			cursor = &lastIssue.Cursor
			lastUpdated := lastIssue.Node.UpdatedAt.Time
			total += cnt
			totalLimit -= cnt
			log.Println(cnt, "fetced", owner, name, lastUpdated)
			if since.After(lastUpdated) {
				break
			}
		}
		if cnt != limit {
			break
		}
	}

	log.Printf("fetched %d pull requests from %s/%s\n", total, owner, name)
	return
}

func (client *GithubInfoV4) GetPullRequestsByNumber(owner, name string, number int) (*PullRequest, error) {
	var query struct {
		Repository struct {
			PullRequest struct {
				PullRequest
			} `graphql:"pullRequest(number: $number)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}
	num := githubv4.Int(number)
	params := map[string]interface{}{
		"number": num,
		"name":   githubv4.String(name),
		"owner":  githubv4.String(owner),
	}
	if err := client.client.Query(context.Background(), &query, params); err != nil {
		return nil, err
	}
	return &query.Repository.PullRequest.PullRequest, nil
}

func (client *GithubInfoV4) GetIssueByNumber(owner, name string, number int) (*IssueNode, error) {
	var query struct {
		Repository struct {
			Issue struct {
				IssueNode
			} `graphql:"issue(number: $number)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}
	num := githubv4.Int(number)
	params := map[string]interface{}{
		"number": num,
		"name":   githubv4.String(name),
		"owner":  githubv4.String(owner),
	}
	if err := client.client.Query(context.Background(), &query, params); err != nil {
		return nil, err
	}
	return &query.Repository.Issue.IssueNode, nil
}
