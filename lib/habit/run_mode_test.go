package habit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMode(t *testing.T) {
	require.True(t, IsModeDebug(ModeDebug))
	require.True(t, IsModeProd(ModeProd))
	require.True(t, IsModeDev(ModeDev))
}
