package extos

import (
	"testing"
)

func TestIsMachineLittleEndian(t *testing.T) {
	t.Run("IsMachineLittleEndian -- ubuntu小端", func(t *testing.T) {
		if !IsMachineLittleEndian() {
			t.Errorf("IsMachineLittleEndian() gotByte3 = %v, want %v", IsMachineLittleEndian(), true)
		}
	})
}

func TestCover(t *testing.T) {
	t.Log(IsWindows())
	t.Log(IsLinux())
	t.Log(IsMac())
}
