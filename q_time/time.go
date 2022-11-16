package q_time

import "time"

// CurrentTime 获取当前时间[12:23:12]
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}
