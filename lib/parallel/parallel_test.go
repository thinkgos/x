package parallel

import (
	"context"
	"errors"
	"fmt"
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
	r, _ := WithContext(context.Background(), maxPar)
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

func TestParallelError(t *testing.T) {
	const totalDo = 5
	r, _ := WithContext(context.Background(), 6)
	for i := 0; i < totalDo; i++ {
		r.Do(func() error {
			return errors.New("error happen")
		})
	}
	err := r.Wait()
	assert.NotNil(t, err)
}

func ExampleNew() {
	p, _ := WithContext(context.Background(), 2)
	p.Do(func() error { return nil })
	err := p.Wait()
	fmt.Println(err)
	// output:
	// <nil>
}
