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

// Sleep pauses the current goroutine for at least the second d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// MSleep pauses the current goroutine for at least the millisecond d.
// A negative or zero duration causes Sleep to return immediately.
func MSleep(t int64) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

// USleep pauses the current goroutine for at least the microsecond d.
// A negative or zero duration causes Sleep to return immediately.
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
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
