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

package track

import (
	"errors"
	"sync"
)

// Context Context's methods may be called by multiple goroutines simultaneously.
type Context interface {
	// Done returns a channel that's closed when work done on behalf of this
	// context should be closed. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously after the close or
	// cancel function called.
	// a Done channel for cancellation.
	Done() <-chan struct{}
	// Err returns the err for the goroutine death provided via Close
	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns ErrDying.
	// If Done is closed with Cancel function, Err returns custom error.
	Err() error
	// Close flags the goroutine as dying.
	Close() error
	// Cancel flags the goroutine as dying with custom error.
	// if error is nil, it will use ErrDying.
	Cancel(err error) error
}

// ErrDying is the error returned by Context.Err when the context is canceled.
var ErrDying = errors.New("track: dying")

// closedChan is a reusable closed channel.
var closedChan = make(chan struct{})

func init() {
	close(closedChan)
}

// Track tracks the lifecycle of a goroutine as alive, dying.
type Track struct {
	mu   sync.Mutex
	done chan struct{}
	err  error
}

// Done a Done channel for cancellation.
func (sf *Track) Done() chan struct{} {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	if sf.done == nil {
		sf.done = make(chan struct{})
	}
	return sf.done
}

// Err returns the err for the goroutine death provided via Close
// If Done is not yet closed, Err returns nil.
// If Done is closed, Err returns ErrDying.
func (sf *Track) Err() error {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	return sf.err
}

// Close flags the goroutine as dying.
func (sf *Track) Close() error {
	return sf.Cancel(nil)
}

// Cancel flags the goroutine as dying with custom error.
// if error is nil, it will use ErrDying.
func (sf *Track) Cancel(err error) error {
	if err == nil {
		err = ErrDying
	}

	sf.mu.Lock()
	defer sf.mu.Unlock()
	if sf.err == nil {
		sf.err = err
		if sf.done == nil {
			sf.done = closedChan
		} else {
			close(sf.done)
		}
	}
	return nil
}
