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

package extos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileModeTime(t *testing.T) {
	tm, err := FileModTime("testdata/a.go")
	require.NoError(t, err)
	assert.True(t, tm > int64(0))

	_, err = FileModTime("files.go")
	require.Error(t, err)
}

func TestFileSize(t *testing.T) {
	size, err := FileSize("testdata/a.go")
	require.NoError(t, err)
	assert.True(t, size >= int64(17)) // windows is 18

	_, err = FileSize("files.go")
	require.Error(t, err)
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile("file.go"))
	assert.False(t, IsFile("testdata"))
	assert.False(t, IsFile("files.go"))
}

func TestIsExist(t *testing.T) {
	t.Run("Pass a file name that exists", func(t *testing.T) {
		assert.True(t, IsExist("file.go"))
	})

	t.Run("Pass a directory name that exists", func(t *testing.T) {
		assert.True(t, IsExist("testdata"))
	})

	t.Run("Pass a directory name that does not exist", func(t *testing.T) {
		assert.False(t, IsExist(".hg"))
	})
}

func TestFileCopy(t *testing.T) {
	src := "testdata/a.go"
	dst := "testdata/a_copy.go"

	err := FileCopy(src, dst)
	require.NoError(t, err)
	defer os.Remove(dst)
	require.True(t, IsExist(dst))
}

func TestWriteFile(t *testing.T) {
	var filename = "testdata/x/y/z.go"
	var testdata = []byte("hello world")
	err := WriteFile(filename, testdata)
	require.NoError(t, err)
	os.RemoveAll("testdata/x")
}

func TestHasPermission(t *testing.T) {
	assert.True(t, HasPermission("file.go"))
}

func BenchmarkIsFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFile("file.go")
	}
}

func BenchmarkIsExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExist("file.go")
	}
}
