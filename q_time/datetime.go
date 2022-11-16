package q_time

import "time"

// CurrentDateTime 获取当前日期+时间[2022-01-01 12:23:12]
func CurrentDateTime() string {
	return time.Now().Format("2006-01-01 15:04:05")
}
