package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		salt := "salt"
		org := "hahaha"
		cpt := new(Simple)

		dst, err := cpt.Hash(org, salt)
		require.Nil(t, err)
		require.NoError(t, cpt.Compare(org, salt, dst))
	})

	t.Run("not correct", func(t *testing.T) {
		salt := "salt"
		org := "hahaha"
		cpt := new(Simple)

		dst, err := cpt.Hash(org, salt)
		require.Nil(t, err)
		require.Error(t, cpt.Compare("invalid", salt, dst))
	})
}

func BenchmarkSimple_Hash(b *testing.B) {
	salt := "salt"
	cpt := new(Simple)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.Hash("hahaha", salt)
	}
}

func BenchmarkSimple_Compare(b *testing.B) {
	salt := "salt"
	org := "hahaha"
	cpt := new(Simple)
	dst, _ := cpt.Hash(org, salt)

	for i := 0; i < b.N; i++ {
		_ = cpt.Compare(org, salt, dst)
	}
}
