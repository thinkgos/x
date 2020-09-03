package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		org := "hahaha"
		cpt := NewSimple("")

		dst, err := cpt.Hash(org)
		require.Nil(t, err)
		require.NoError(t, cpt.Compare(org, dst))
	})

	t.Run("not correct", func(t *testing.T) {
		org := "hahaha"
		cpt := NewSimple("")

		dst, err := cpt.Hash(org)
		require.Nil(t, err)
		require.Error(t, cpt.Compare("invalid", dst))
	})
}

func BenchmarkSimple_Hash(b *testing.B) {
	cpt := NewSimple("")

	for i := 0; i < b.N; i++ {
		_, _ = cpt.Hash("hahaha")
	}
}

func BenchmarkSimple_Compare(b *testing.B) {
	org := "hahaha"
	cpt := NewSimple("")
	dst, _ := cpt.Hash(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.Compare(org, dst)
	}
}
