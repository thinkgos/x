// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ciol 实现 net.conn 网络io限速器接口
package ciol

import (
	"context"
	"net"
	"time"

	"golang.org/x/time/rate"
)

// Conn limiter conn
type Conn struct {
	net.Conn
	rLimiter *rate.Limiter
	wLimiter *rate.Limiter
	ctx      context.Context
}

// New new a rate limit (bytes/sec) with option to the Conn read and write.
// if not set,it will not any limit
func New(c net.Conn, opts ...Options) *Conn {
	s := &Conn{
		Conn: c,
		ctx:  context.Background(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	n, err := sf.Conn.Read(p)
	if err != nil || sf.rLimiter == nil {
		return n, err
	}
	return n, sf.rLimiter.WaitN(sf.ctx, n)
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	n, err := sf.Conn.Write(p)
	if err != nil || sf.wLimiter == nil {
		return n, err
	}
	return n, sf.wLimiter.WaitN(sf.ctx, n)
}

// Close close the Conn
func (sf *Conn) Close() (err error) {
	return sf.Conn.Close()
}

// ReadLimit returns the maximum overall event read rate.
func (sf *Conn) ReadLimit() rate.Limit {
	if sf.rLimiter != nil {
		return sf.rLimiter.Limit()
	}
	return 0
}

// SetReadLimit sets a new read Limit for the limiter.
func (sf *Conn) SetReadLimit(newLimit rate.Limit) {
	if sf.rLimiter != nil {
		sf.rLimiter.SetLimit(newLimit)
	}
}

// SetReadLimitAt sets a new read Limit for the limiter.
func (sf *Conn) SetReadLimitAt(now time.Time, newLimit rate.Limit) {
	if sf.rLimiter != nil {
		sf.rLimiter.SetLimitAt(now, newLimit)
	}
}

// SetReadBurst sets a new read burst size for the limiter.
func (sf *Conn) SetReadBurst(newBurst int) {
	if sf.rLimiter != nil {
		sf.rLimiter.SetBurst(newBurst)
	}
}

// SetReadBurstAt sets a new read read size for the limiter.
func (sf *Conn) SetReadBurstAt(now time.Time, newBurst int) {
	if sf.rLimiter != nil {
		sf.rLimiter.SetBurstAt(now, newBurst)
	}
}

// WriteLimit returns the maximum overall event write rate.
func (sf *Conn) WriteLimit() rate.Limit {
	if sf.wLimiter != nil {
		return sf.wLimiter.Limit()
	}
	return 0
}

// SetWriteLimit sets a new write Limit for the limiter.
func (sf *Conn) SetWriteLimit(newLimit rate.Limit) {
	if sf.wLimiter != nil {
		sf.wLimiter.SetLimit(newLimit)
	}
}

// SetWriteLimitAt sets a new write Limit for the limiter.
func (sf *Conn) SetWriteLimitAt(now time.Time, newLimit rate.Limit) {
	if sf.wLimiter != nil {
		sf.wLimiter.SetLimitAt(now, newLimit)
	}
}

// SetWriteBurst sets a new read write size for the limiter.
func (sf *Conn) SetWriteBurst(newBurst int) {
	if sf.wLimiter != nil {
		sf.wLimiter.SetBurst(newBurst)
	}
}

// SetWriteBurstAt sets a new read write size for the limiter.
func (sf *Conn) SetWriteBurstAt(now time.Time, newBurst int) {
	if sf.wLimiter != nil {
		sf.wLimiter.SetBurstAt(now, newBurst)
	}
}
