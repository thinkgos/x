package password

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBCrypt(t *testing.T) {
	salt := "salt"
	org := "hahaha"
	cpt := new(BCrypt)

	dst, err := cpt.Hash(org, salt)
	require.Nil(t, err)
	require.Nil(t, cpt.Compare(org, salt, dst))
}

func BenchmarkBCrypt_Hash(b *testing.B) {
	salt := "salt"
	cpt := new(BCrypt)

	for i := 0; i < b.N; i++ {
		_, _ = cpt.Hash("hahaha", salt)
	}
}

func BenchmarkBCrypt_Compare(b *testing.B) {
	salt := "salt"
	org := "hahaha"
	cpt := new(BCrypt)
	dst, _ := cpt.Hash(org, salt)

	for i := 0; i < b.N; i++ {
		_ = cpt.Compare(org, salt, dst)
	}
}
