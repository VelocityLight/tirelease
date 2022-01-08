package service

import (
	"fmt"
	"time"

	"tirelease/internal/entity"
	"tirelease/internal/repository"
)

func IssueAffectOperate(updateOption *entity.IssueAffectUpdateOption) error {
	// Params Check
	option := &entity.IssueAffectOption{}
	if updateOption.ID != 0 {
		option.ID = updateOption.ID
	} else if updateOption.IssueID != "" && updateOption.AffectVersion != "" {
		option.IssueID = updateOption.IssueID
		option.AffectVersion = updateOption.AffectVersion
	} else {
		return fmt.Errorf("update option params key is empty: %+v failed", updateOption)
	}
	if updateOption.AffectResult == "" {
		return fmt.Errorf("update option params values is empty: %+v failed", updateOption)
	}

	// Select
	issueAffects, err := repository.SelectIssueAffect(option)
	if nil != err {
		return err
	}
	if len(*issueAffects) != 1 {
		return fmt.Errorf("select issue affect is empty or many: %+v failed", *issueAffects)
	}

	// Update issue_affect
	(*issueAffects)[0].AffectResult = updateOption.AffectResult
	err = repository.UpdateIssueAffect(&(*issueAffects)[0])
	if nil != err {
		return err
	}

	// Insert version_triage
	if (*issueAffects)[0].AffectResult == entity.AffectResultResultYes {
		versionTriage := &entity.VersionTriage{
			VersionName:  (*issueAffects)[0].AffectVersion,
			IssueID:      (*issueAffects)[0].IssueID,
			TriageResult: entity.VersionTriageResultUnKnown,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		err = repository.CreateVersionTriage(versionTriage)
		if nil != err {
			return err
		}
	}

	return nil
}
