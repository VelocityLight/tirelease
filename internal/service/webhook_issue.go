package service

// import (
// 	"time"

// 	"tirelease/internal/entity"
// 	"tirelease/internal/repository"
// )

// // CronJob
// func InitIssueAndRelationFirstTime() {
// 	err := UpdateIssueAndRelationInfoByTime(&entity.RepoOption{}, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func UpdateIssueAndRelationInfoByTime(repoOption *entity.RepoOption, time *time.Time) error {
// 	// Get Repos
// 	repos, err := repository.SelectRepo(repoOption)
// 	if err != nil {
// 		return err
// 	}

// 	// Save Issues Info
// 	for _, repo := range *repos {
// 		issues, err := GetIssuesByTimeFromV3(repo.Owner, repo.Repo, time)
// 		if err != nil {
// 			return err
// 		}

// 		for _, issue := range issues {
// 			issueRelation, err := ComposeTriageRelationInfoByIssue(issue)
// 			if err != nil {
// 				return err
// 			}
// 			err = SaveIssueRelationInfo(issueRelation)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// // Git Webhook

// // Inner Service
