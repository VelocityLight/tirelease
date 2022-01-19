package cron

import (
	"fmt"

	"tirelease/commons/cron"
)

func DemoCron() {
	// Cron 表达式及功能方法
	cron.Create("*/5 * * * * *", func() { fmt.Println("Every Five Seconds") })
}
