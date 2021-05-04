package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSCrypt(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		org := "hahaha"
		cpt := new(SCrypt)

		dst, err := cpt.GenerateFromPassword(org)
		t.Log(dst)
		require.Nil(t, err)
		require.NoError(t, cpt.CompareHashAndPassword(dst, org))
	})

	t.Run("not correct", func(t *testing.T) {
		org := "hahaha"
		cpt := new(SCrypt)

		dst, err := cpt.GenerateFromPassword(org)
		require.Nil(t, err)
		require.Error(t, cpt.CompareHashAndPassword(dst, "invalid"))
	})
}

func BenchmarkSCrypt_GenerateFromPassword(b *testing.B) {
	cpt := new(SCrypt)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.GenerateFromPassword("hahaha")
	}
}

func BenchmarkSCrypt_CompareHashAndPassword(b *testing.B) {
	org := "hahaha"
	cpt := new(SCrypt)
	dst, _ := cpt.GenerateFromPassword(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.CompareHashAndPassword(dst, org)
	}
}
