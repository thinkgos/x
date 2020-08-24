package ternary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTernary(t *testing.T) {
	t.Run("interface", func(t *testing.T) {
		require.Equal(t, 1, If(true, 1, "string"))
		require.Equal(t, "string", If(false, 1, "string"))
	})
	t.Run("int", func(t *testing.T) {
		require.Equal(t, 1, IfInt(true, 1, 2))
		require.Equal(t, 2, IfInt(false, 1, 2))
	})
	t.Run("int64", func(t *testing.T) {
		require.Equal(t, int64(1), IfInt64(true, 1, 2))
		require.Equal(t, int64(2), IfInt64(false, 1, 2))
	})
	t.Run("float", func(t *testing.T) {
		require.Equal(t, float64(1.1), IfFloat(true, 1.1, 2.1))
		require.Equal(t, float64(2.1), IfFloat(false, 1.1, 2.1))
	})
	t.Run("string", func(t *testing.T) {
		require.Equal(t, "true", IfString(true, "true", "false"))
		require.Equal(t, "false", IfString(false, "true", "false"))
	})
}
