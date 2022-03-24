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
	sql := "select * from issue where 1=1" + IssueWhere(option) + IssueOrderBy(option) + IssueLimit(option)

	// 查询
	var issues []entity.Issue
	if err := database.DBConn.RawWrapper(sql, option).Find(&issues).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("select issue by raw failed, option: %+v", option))
	}
	return &issues, nil
}

func CountIssueRaw(option *entity.IssueOption) (int64, error) {
	sql := "select count(*) from issue where 1=1" + IssueWhere(option) + IssueOrderBy(option) + IssueLimit(option)

	var count int64
	if err := database.DBConn.RawWrapper(sql, option).Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("count issue by raw failed, option: %+v", option))
	}
	return count, nil
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

func IssueWhere(option *entity.IssueOption) string {
	sql := ""

	if option.ID != 0 {
		sql += " and id = @ID"
	}
	if option.IssueID != "" {
		sql += " and issue_id = @IssueID"
	}
	if option.Number != 0 {
		sql += " and number = @Number"
	}
	if option.State != "" {
		sql += " and state = @State"
	}
	if option.Owner != "" {
		sql += " and owner = @Owner"
	}
	if option.Repo != "" {
		sql += " and repo = @Repo"
	}
	if option.SeverityLabel != "" {
		sql += " and severity_label = @SeverityLabel"
	}
	if option.TypeLabel != "" {
		sql += " and type_label = @TypeLabel"
	}
	if option.IssueIDs != nil && len(option.IssueIDs) > 0 {
		sql += " and issue_id in @IssueIDs"
	}
	if option.SeverityLabels != nil && len(option.SeverityLabels) > 0 {
		sql += " and severity_label in @SeverityLabels"
	}
	if option.NotSeverityLabels != nil && len(option.NotSeverityLabels) > 0 {
		sql += " and severity_label not in @NotSeverityLabels"
	}

	return sql
}

func IssueOrderBy(option *entity.IssueOption) string {
	sql := ""

	if option.OrderBy != "" {
		sql += " order by " + option.OrderBy
	}
	if option.Order != "" {
		sql += " " + option.Order
	}

	return sql
}

func IssueLimit(option *entity.IssueOption) string {
	sql := ""

	if option.Page != 0 && option.PerPage != 0 {
		option.ListOption.CalcOffset()
		sql += " limit @Offset,@PerPage"
	}

	return sql
}
