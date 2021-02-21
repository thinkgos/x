package extjson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	var ret struct {
		ID string `json:"id"`
	}
	err := UnmarshalFromString(`{"id":"123"}`, &ret)
	require.NoError(t, err)
	require.Equal(t, "123", ret.ID)

	s, err := MarshalToString(ret)
	require.NoError(t, err)
	require.Equal(t, `{"id":"123"}`, s)
}
