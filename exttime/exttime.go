package exttime

import "time"

var setupTime = time.Now()

// Millisecond time.Time 转为 毫秒
func Millisecond(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Millisecond)
}

// Microsecond time.Time 转为 微秒
func Microsecond(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Microsecond)
}

// Time 毫秒转time.Time
func Time(msec int64) time.Time {
	return time.Unix(msec/1000, (msec%1000)*int64(time.Millisecond))
}

// ServiceStartupTime 服务启动时间
func ServiceStartupTime() time.Time {
	return setupTime
}

// ServiceElapseTime 服务启动了多少时间(nanosecond)
func ServiceElapseTime() time.Duration {
	return time.Since(setupTime)
}

// ServiceUptime 服务启动了多少second
func ServiceUptime() int64 {
	return int64(time.Since(setupTime) / time.Second)
}
