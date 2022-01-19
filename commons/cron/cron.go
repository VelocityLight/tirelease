// Tool Url: https://github.com/robfig/cron

package cron

import (
	"github.com/robfig/cron"
)

// ┌─────────────second 范围 (0 - 60)
// │ ┌───────────── min (0 - 59)
// │ │ ┌────────────── hour (0 - 23)
// │ │ │ ┌─────────────── day of month (1 - 31)
// │ │ │ │ ┌──────────────── month (1 - 12)
// │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
// │ │ │ │ │ │                  Saturday)
// │ │ │ │ │ │
// │ │ │ │ │ │
//  *  *  *  *  *  *
//
// 匹配符号如下:
// 星号(*): 表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
// 斜线(/): 表示增长间隔，如第2个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后 每隔 15 分钟执行一次（即 3（3+0*15）、18（3+1*15）、33（3+2*15）、48（3+3*15） 这些时间点执行），这里也可以表示为：3/15
// 逗号(,): 用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
// 连字号(-): 表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 之间每个小时（包括9和17）
// 问号(?): 只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *

// Create Function
func Create(expression string, f func()) *cron.Cron {
	c := cron.New()
	c.AddFunc(expression, f)
	c.Start()
	return c
}

// Stop Function
func Stop(c *cron.Cron) {
	defer c.Stop()
}
