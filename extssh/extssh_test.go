package extssh

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var passPhrase = []byte("passPhrase")

func TestLoadPrivateKeyFile2AuthMethod(t *testing.T) {
	_, err := LoadPrivateKey2AuthMethod("./testdata/id_rsa")
	require.NoError(t, err)

	_, err = LoadPrivateKey2AuthMethod("./testdata/id_rsa_passPhrase", passPhrase)
	require.NoError(t, err)

	_, err = LoadPrivateKey2AuthMethod("./testdata/nofile")
	require.Error(t, err)
}
