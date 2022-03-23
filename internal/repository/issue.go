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

// func CreateIssue(issue *entity.Issue) error {
// 	// 加工
// 	serializeIssue(issue)

// 	// 存储
// 	if err := database.DBConn.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&issue).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("create issue: %+v failed", issue))
// 	}
// 	return nil
// }

func SelectIssue(option *entity.IssueOption) (*[]entity.Issue, error) {
	// 查询
	var issues []entity.Issue
	if err := database.DBConn.DB.Where(option).Order("updated_at desc").Find(&issues).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find issue: %+v failed", option))
	}

	// 加工
	for i := 0; i < len(issues); i++ {
		unSerializeIssue(&issues[i])
	}
	return &issues, nil
}

func SelectIssueRaw(option *entity.IssueOption) (*[]entity.Issue, error) {
	// 查询
	var issues []entity.Issue
	sql, equal := SelectIssueSQL("select * from issue where 1=1", option)
	if equal {
		if err := database.DBConn.DB.Raw(sql).Order("updated_at desc").Find(&issues).Error; err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("select issue by raw failed, option: %+v", option))
		}
	} else {
		if err := database.DBConn.DB.Raw(sql, option).Order("updated_at desc").Find(&issues).Error; err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("select issue by raw failed, option: %+v", option))
		}
	}
	return &issues, nil
}

func SelectIssueUnique(option *entity.IssueOption) (*entity.Issue, error) {
	// 查询
	issues, err := SelectIssue(option)
	if err != nil {
		return nil, err
	}

	// 校验
	if len(*issues) == 0 {
		return nil, errors.New(fmt.Sprintf("issue not found: %+v", option))
	}
	if len(*issues) > 1 {
		return nil, errors.New(fmt.Sprintf("more than one issue found: %+v", option))
	}
	return &((*issues)[0]), nil
}

// func DeleteIssue(issue *entity.Issue) error {
// 	if err := database.DBConn.DB.Delete(issue).Error; err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("delete issue: %+v failed", issue))
// 	}
// 	return nil
// }

func CreateOrUpdateIssue(issue *entity.Issue) error {
	// 加工
	serializeIssue(issue)

	// 存储
	if err := database.DBConn.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Omit("Labels", "Assignees").Create(&issue).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create or update issue: %+v failed", issue))
	}
	return nil
}

// 序列化和反序列化
func serializeIssue(issue *entity.Issue) {
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
	if issue.AssigneesString != "" {
		var assignees []github.User
		json.Unmarshal([]byte(issue.AssigneesString), &assignees)
		issue.Assignees = &assignees
	}
	if issue.LabelsString != "" {
		var labels []github.Label
		json.Unmarshal([]byte(issue.LabelsString), &labels)
		issue.Labels = &labels
	}
}

func SelectIssueSQL(sql string, option *entity.IssueOption) (string, bool) {
	newSql := string(sql)
	if option.ID != 0 {
		sql += " and id = @ID"
	}
	if option.IssueID != "" {
		sql += " and issue_id = @IssueID"
	}
	if option.IssueIDs != nil && len(option.IssueIDs) > 0 {
		sql += " and issue_id in @IssueIDs"
	}
	return sql, newSql == sql
}
