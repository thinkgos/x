// The parallel package provides a way of running functions
// concurrently while limiting the maximum number running at once.
// same as errgroup.Group but with a limiter number
package parallel

import (
	"context"
	"sync"
)

// Parallel represents a number of functions running concurrently.
type Parallel struct {
	limiter chan struct{}
	cancel  func()

	wg sync.WaitGroup

	once sync.Once
	err  error
}

// WithContext returns a new parallel instance and an associated Context derived from ctx..  It will run up to maxPar
// functions concurrently.
func WithContext(ctx context.Context, maxPar int) (*Parallel, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Parallel{
		limiter: make(chan struct{}, maxPar),
		cancel:  cancel,
	}, ctx
}

// Do requests that r run f concurrently.  If there are already the maximum
// number of functions running concurrently, it will block until one of
// them has completed. Do may itself be called concurrently.
func (r *Parallel) Do(f func() error) {
	r.wg.Add(1)
	r.limiter <- struct{}{}
	go func() {
		defer func() {
			r.wg.Done()
			<-r.limiter
		}()
		if err := f(); err != nil {
			r.once.Do(func() {
				r.err = err
				if r.cancel != nil {
					r.cancel()
				}
			})
		}
	}()
}

// Wait marks the parallel instance as complete and waits for all the
// functions to complete.  If any errors were encountered, it returns an
// Errors value describing all the errors in arbitrary order.
func (r *Parallel) Wait() error {
	r.wg.Wait()
	if r.cancel != nil {
		r.cancel()
	}
	return r.err
}
