package extos

import (
	"os"
	"path/filepath"
)

// Executable 和os.Executable() 一致,但是获取执行程序的所在路径和文件名
func Executable() (dir, file string, err error) {
	path, err := os.Executable()
	if err != nil {
		return "", "", err
	}
	return filepath.Dir(path), filepath.Base(path), nil
}
