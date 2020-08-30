package extnet

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testError struct {
	err string
}

func (e *testError) Error() string   { return e.err }
func (e *testError) Timeout() bool   { return true }
func (e *testError) Temporary() bool { return true }

func TestErr(t *testing.T) {
	assert.True(t, IsErrClosed(errors.New("use of closed network connection")))
	assert.False(t, IsErrClosed(nil))

	assert.True(t, IsErrTimeout(&testError{"timeout"}))
	assert.False(t, IsErrTimeout(nil))

	assert.True(t, IsErrTemporary(&testError{"temporary"}))
	assert.False(t, IsErrTemporary(nil))

	assert.True(t, IsErrRefused(errors.New("connection refused")))
	assert.False(t, IsErrRefused(nil))

	assert.True(t, IsErrDeadline(errors.New("i/o deadline reached")))
	assert.False(t, IsErrDeadline(nil))

	assert.True(t, IsErrSocketNotConnected(errors.New("socket is not connected")))
	assert.False(t, IsErrSocketNotConnected(nil))
}
