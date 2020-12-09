// Copyright 2013 com authors
// authors: Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package extos base tool
package extos

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// FileModTime returns file modified time and possible error.
func FileModTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// FileSize returns file size in bytes and possible error.
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// IsDir returns true if given path is a dir,
// or returns false when it's a file or does not exist.
func IsDir(filePath string) bool {
	f, err := os.Stat(filePath)
	return err == nil && f.IsDir()
}

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, err := os.Stat(filePath)
	return err == nil && !f.IsDir()
}

// FileMode returns file mode and possible error.
func FileMode(name string) (os.FileMode, error) {
	fInfo, err := os.Lstat(name)
	if err != nil {
		return 0, err
	}
	return fInfo.Mode(), nil
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(paths string) bool {
	_, err := os.Stat(paths)
	return err == nil || os.IsExist(err)
}

// HasPermission returns a boolean indicating whether that permission is allowed.
func HasPermission(name string) bool {
	_, err := os.Stat(name)
	return !os.IsPermission(err)
}

// FileCopy copies file from source to target path.
func FileCopy(src, dest string) error {
	// Gather file information to set back later.
	si, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Handle symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// NOTE: os.Chmod and os.Chtimes don't recoganize symbolic link,
		// which will lead "no such file or directory" error.
		return os.Symlink(target, dest)
	}

	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}

	// Set back file information.
	if err := os.Chtimes(dest, si.ModTime(), si.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dest, si.Mode())
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it
// and its upper level paths.
func WriteFile(filename string, data []byte) error {
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0655)
}

// Filepaths returns all root dir (contain sub dir) file full path
func Filepaths(root string) ([]string, error) {
	var result = make([]string, 0)

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, err
}

// // IsExecutable 是否可执行文件.
// TODO: bug on wondows
// func IsExecutable(name string) bool {
// 	info, err := os.Stat(name)
// 	return err == nil && info.Mode().IsRegular() && (info.Mode()&0111) != 0
// }
//
// // IsLink 是否链接文件(且存在).
// TODO: bug on wondows
// func IsLink(name string) bool {
// 	f, err := os.Lstat(name)
// 	return err == nil && f.Mode()&os.ModeSymlink == os.ModeSymlink
// }
