package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBCrypt(t *testing.T) {
	org := "hahaha"
	cpt := new(BCrypt)

	dst, err := cpt.GenerateFromPassword(org)
	t.Log(dst)
	require.Nil(t, err)
	require.Nil(t, cpt.CompareHashAndPassword(dst, org))
}

func BenchmarkBCrypt_GenerateFromPassword(b *testing.B) {
	cpt := new(BCrypt)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.GenerateFromPassword("hahaha")
	}
}

func BenchmarkBCrypt_CompareHashAndPassword(b *testing.B) {
	org := "hahaha"
	cpt := new(BCrypt)
	dst, _ := cpt.GenerateFromPassword(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.CompareHashAndPassword(dst, org)
	}
}
