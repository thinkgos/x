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

func TestImage(t *testing.T) {
	want := "helloworld"
	str := EncodeToBase64("png", []byte(want))
	require.True(t, strings.HasPrefix(str, "data:image/png;base64,"))

	imgType, imgValue, err := DecodeBase64(str)
	require.NoError(t, err)
	require.Equal(t, "image/png", imgType)
	require.Equal(t, []byte(want), imgValue)

	str1 := EncodeToBase64("image/png", []byte(want))
	require.True(t, strings.HasPrefix(str1, "data:image/png;base64,"))

	imgType1, imgValue1, err := DecodeBase64(str)
	require.NoError(t, err)
	require.Equal(t, "image/png", imgType1)
	require.Equal(t, []byte(want), imgValue1)

	_, _, err = DecodeBase64("invalid_image")
	require.Error(t, err)
}

func BenchmarkEncodeToBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeToBase64("png", []byte("helloworld"))
	}
}

func BenchmarkDecodeBase64(b *testing.B) {
	s := EncodeToBase64("png", []byte("helloworld"))
	for i := 0; i < b.N; i++ {
		_, _, _ = DecodeBase64(s)
	}
}
