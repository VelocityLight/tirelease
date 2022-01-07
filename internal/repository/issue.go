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

func CreateOrUpdateIssue(issue *entity.Issue) error {
	// 加工
	serializeIssue(issue)

	// 存储
	if err := database.DBConn.DB.Clauses(
		clause.OnConflict{UpdateAll: true}).Omit("Labels", "Assignee", "Assignees").Create(&issue).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create issue: %+v failed", issue))
	}
	return nil
}

func SelectIssue(option *entity.IssueOption) (*[]entity.Issue, error) {
	// 查询
	var issues []entity.Issue
	if err := database.DBConn.DB.Find(&issues).Where(option).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find issue: %+v failed", option))
	}

	// 加工
	for i := 0; i < len(issues); i++ {
		unSerializeIssue(&issues[i])
	}
	return &issues, nil
}

func DeleteIssue(issue *entity.Issue) error {
	if err := database.DBConn.DB.Delete(issue).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("delete issue: %+v failed", issue))
	}
	return nil
}

// 序列化和反序列化
func serializeIssue(issue *entity.Issue) {
	if nil != issue.Assignee {
		assigneeString, _ := json.Marshal(issue.Assignee)
		issue.AssigneeString = string(assigneeString)
	}
	if nil != issue.Assignees {
		assigneesString, _ := json.Marshal(issue.Assignees)
		issue.AssigneesString = string(assigneesString)
	}
	if nil != issue.Labels {
		labelsString, _ := json.Marshal(issue.Labels)
		issue.LabelsString = string(labelsString)
	}
}

func unSerializeIssue(issue *entity.Issue) {
	if issue.AssigneeString != "" {
		var assignee github.User
		json.Unmarshal([]byte(issue.AssigneeString), &assignee)
		issue.Assignee = &assignee
	}
	if issue.AssigneesString != "" {
		var assignees []github.User
		json.Unmarshal([]byte(issue.AssigneeString), &assignees)
		issue.Assignees = &assignees
	}
	if issue.LabelsString != "" {
		var labels []github.Label
		json.Unmarshal([]byte(issue.LabelsString), &labels)
		issue.Labels = &labels
	}
}
