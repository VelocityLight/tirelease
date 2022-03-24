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
		BeforeHours: -25,
		Batch:       20,
		Total:       500,
	}
	cron.Create("0 0 */1 * * ?", func() { service.CronRefreshIssuesV4(params) })
}
