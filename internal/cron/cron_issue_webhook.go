package cron

import (
	"tirelease/commons/cron"
	"tirelease/internal/service"
)

func IssueWebHook() {
	// Cron 表达式及功能方法
	cron.Create("* */1 * * * *", func() { service.CronRefreshIssuesV4() })
}
