package parallel

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParallelMaxPar(t *testing.T) {
	const (
		totalDo = 10
		maxPar  = 3
	)
	var mu sync.Mutex
	max := 0
	n := 0
	total := 0
	r := New(maxPar)
	for i := 0; i < totalDo; i++ {
		r.Do(func() error {
			mu.Lock()
			total++
			n++
			if n > max {
				max = n
			}
			mu.Unlock()
			time.Sleep(time.Millisecond * 100)
			mu.Lock()
			n--
			mu.Unlock()
			return nil
		})
	}
	err := r.Wait()
	assert.Nil(t, err)
	assert.Equalf(t, 0, n, "%d functions still running", n)
	assert.Equalf(t, totalDo, total, "all functions not executed; want %d got %d", totalDo, total)
	assert.Equalf(t, maxPar, max, "wrong number of do's ran at once; want %d got %d", maxPar, max)
}

type intError int

func (intError) Error() string {
	return "error"
}

func TestParallelError(t *testing.T) {
	const (
		totalDo = 10
		errDo   = 5
	)
	r := New(6)
	for i := 0; i < totalDo; i++ {
		i := i
		if i >= errDo {
			r.Do(func() error {
				return intError(i)
			})
		} else {
			r.Do(func() error {
				return nil
			})
		}
	}
	err := r.Wait()
	assert.NotNil(t, err)
	errs := err.(Errors)
	assert.Equalf(t, totalDo-errDo, len(errs), "wrong error count; want %d got %d", len(errs), totalDo-errDo)
	ints := make([]int, len(errs))
	for i, err := range errs {
		ints[i] = int(err.(intError))
	}
	sort.Ints(ints)
	for i, n := range ints {
		assert.Equalf(t, i+errDo, n, "unexpected error value; want %d got %d", i+errDo, n)
	}
}

func ExampleNew() {
	p := New(2)
	p.Do(
		func() error { return nil },
		func() error { return nil },
		func() error { return nil })
	err := p.Wait()
	fmt.Println(err)
	// output:
	// <nil>
}
