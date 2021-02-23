package track

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrack(t *testing.T) {
	goo := &Track{}

	require.NoError(t, goo.Err())
	done := goo.Done()
	select {
	case <-done:
		require.Fail(t, "done chan must not select")
	default:
	}

	require.NoError(t, goo.Close())
	require.Equal(t, ErrDying, goo.Err())

	_, ok := <-done
	require.False(t, ok)
}

func TestTrack_no_done_and_custom_cancel(t *testing.T) {
	goo := &Track{}
	err := errors.New("custom error")
	require.NoError(t, goo.Err())
	require.NoError(t, goo.Cancel(err))
	require.Equal(t, err, goo.Err())
}
