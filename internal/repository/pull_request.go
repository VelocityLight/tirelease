package repository

import (
	"encoding/json"
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/google/go-github/v41/github"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

// func CreatePullRequest(pullRequest *entity.PullRequest) error {
// 	// 加工
// 	serializePullRequest(pullRequest)

// 	// 存储
// 	if err := database.DBConn.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&pullRequest).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("create pull request: %+v failed", pullRequest))
// 	}
// 	return nil
// }

func SelectPullRequest(option *entity.PullRequestOption) (*[]entity.PullRequest, error) {
	// 查询
	var prs []entity.PullRequest
	if err := database.DBConn.DB.Where(option).Order("updated_at desc").Find(&prs).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find pull request: %+v failed", option))
	}

	// 加工
	for i := 0; i < len(prs); i++ {
		unSerializePullRequest(&prs[i])
	}
	return &prs, nil
}

func SelectPullRequestUnique(option *entity.PullRequestOption) (*entity.PullRequest, error) {
	var pr entity.PullRequest
	if err := database.DBConn.DB.Where(option).First(&pr).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find pull request unique: %+v failed", option))
	}

	if pr.PullRequestID == "" {
		return nil, errors.New(fmt.Sprintf("pull request not found: %+v", option))
	}

	// 加工
	unSerializePullRequest(&pr)
	return &pr, nil
}

// func DeletePullRequest(pullRequest *entity.PullRequest) error {
// 	if err := database.DBConn.DB.Delete(pullRequest).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("delete pull request: %+v failed", pullRequest))
// 	}
// 	return nil
// }

func CreateOrUpdatePullRequest(pullRequest *entity.PullRequest) error {
	// 加工
	serializePullRequest(pullRequest)

	// 存储
	if err := database.DBConn.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Omit("Labels", "Assignees", "RequestedReviewers").Create(&pullRequest).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create or update pull request: %+v failed", pullRequest))
	}
	return nil
}

// 序列化和反序列化
func serializePullRequest(pullRequest *entity.PullRequest) {
	if nil != pullRequest.Assignees {
		assigneesString, _ := json.Marshal(pullRequest.Assignees)
		pullRequest.AssigneesString = string(assigneesString)
	}
	if nil != pullRequest.Labels {
		labelsString, _ := json.Marshal(pullRequest.Labels)
		pullRequest.LabelsString = string(labelsString)
	}
	if nil != pullRequest.RequestedReviewers {
		requestedReviewersString, _ := json.Marshal(pullRequest.RequestedReviewers)
		pullRequest.RequestedReviewersString = string(requestedReviewersString)
	}
}

func unSerializePullRequest(pullRequest *entity.PullRequest) {
	if pullRequest.AssigneesString != "" {
		var assignees []github.User
		json.Unmarshal([]byte(pullRequest.AssigneesString), &assignees)
		pullRequest.Assignees = &assignees
	}
	if pullRequest.LabelsString != "" {
		var labels []github.Label
		json.Unmarshal([]byte(pullRequest.LabelsString), &labels)
		pullRequest.Labels = &labels
	}
	if pullRequest.RequestedReviewersString != "" {
		var requestedReviewers []github.User
		json.Unmarshal([]byte(pullRequest.RequestedReviewersString), &requestedReviewers)
		pullRequest.RequestedReviewers = &requestedReviewers
	}
}
