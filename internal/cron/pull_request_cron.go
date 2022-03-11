package cron

import (
	"tirelease/commons/cron"
	"tirelease/internal/service"
)

func PullRequestCron() {
	// Cron 表达式及功能方法
	cron.Create("* */1 * * * *", func() { service.CronRefreshPullRequestV4() })
}
