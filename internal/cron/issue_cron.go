package cron

import (
	"tirelease/commons/cron"
	"tirelease/internal/entity"
	"tirelease/internal/repository"
	"tirelease/internal/service"
)

func IssueCron() {
	// Cron 表达式及功能方法
	repoOption := &entity.RepoOption{}
	repos, err := repository.SelectRepo(repoOption)
	if err != nil {
		return
	}
	params := &service.RefreshIssueParams{
		Repos:       repos,
		BeforeHours: -360,
		Batch:       20,
		Total:       3000,
	}
	cron.Create("* */1 * * * *", func() { service.CronRefreshIssuesV4(params) })
}
