package habit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/thinkgos/x/extos"
)

// WritePidFile 写pid到文件中,默认写到执行文件目录下,并以执行文件名.pid命名
// 可选指定相应目录
func WritePidFile(dir ...string) error {
	path, filename, err := extos.Executable()
	if err != nil {
		return err
	}
	if len(dir) > 0 {
		path = dir[0]
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	pidFilename := filepath.Join(path, filename+".pid")
	pid := os.Getpid()
	return ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

// RemovePidFile 删除默认写到执行文件目录下,并以执行文件名.pid命名的文件.
// 可选指定相应目录
func RemovePidFile(dir ...string) error {
	path, filename, err := extos.Executable()
	if err != nil {
		return err
	}
	if len(dir) > 0 {
		path = dir[0]
	}
	return os.Remove(filepath.Join(path, filename+".pid"))
}
