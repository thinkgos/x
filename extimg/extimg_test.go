package extimg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetType(t *testing.T) {
	imgType, err := GetType("golang.png")
	require.NoError(t, err)
	require.True(t, isImgTypeExt(imgType))
}

func isImgTypeExt(p string) bool {
	for _, ext := range GetExts() {
		if strings.Contains(ext, p[6:]) {
			return true
		}
	}
	return false
}
