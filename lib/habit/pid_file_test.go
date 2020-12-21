package habit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPIDFile(t *testing.T) {
	dir := "/tmp/pidfile"
	require.NoError(t, WritePidFile(dir))
	require.NoError(t, RemovePidFile(dir))
}
