package extime

import "time"

var setupTime = time.Now()

// ServiceStartupTime 服务启动时间
func ServiceStartupTime() time.Time {
	return setupTime
}

// ServiceElapseTime 服务启动了多少时间
func ServiceElapseTime() time.Duration {
	return time.Since(setupTime)
}

// ServiceUptime 服务启动了多少second
func ServiceUptime() int64 {
	return int64(time.Since(setupTime) / time.Second)
}
