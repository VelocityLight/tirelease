package git

import (
	"strings"
)

// Distinguish issue or pull request method
func IsIssue(url string) bool {
	if url == "" {
		return false
	}
	return strings.Contains(url, "/issues/")
}

func IsPullRequest(url string) bool {
	if url == "" {
		return false
	}
	return strings.Contains(url, "/pull/")
}
