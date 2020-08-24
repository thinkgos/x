package ciol

import (
	"time"

	"golang.org/x/time/rate"
)

// MaxBurst 默认读写最大容量限制
const MaxBurst = 1000 * 1000 * 1000

// Options Conn options
type Options func(*Conn)

// WithReadLimiter 读限速
func WithReadLimiter(bytesPerSec rate.Limit, bursts ...int) Options {
	return func(c *Conn) {
		burst := MaxBurst
		if len(bursts) > 0 {
			burst = bursts[0]
		}
		c.rLimiter = rate.NewLimiter(bytesPerSec, burst)
		c.rLimiter.AllowN(time.Now(), burst) // spend initial burst
	}
}

// WithWriteLimiter 写限速
func WithWriteLimiter(bytesPerSec rate.Limit, bursts ...int) Options {
	return func(c *Conn) {
		burst := MaxBurst
		if len(bursts) > 0 {
			burst = bursts[0]
		}
		c.wLimiter = rate.NewLimiter(bytesPerSec, burst)
		c.wLimiter.AllowN(time.Now(), burst) // spend initial burst
	}
}
