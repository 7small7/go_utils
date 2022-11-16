package q_time

import "time"

// CurrentDate 获取当前日期[2022-01-01]
func CurrentDate() string {
	return time.Now().Format("2006-01-01")
}
