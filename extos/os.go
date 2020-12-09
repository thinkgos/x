// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extos base tool
package extos

import (
	"runtime"
	"unsafe"
)

// IsMachineLittleEndian 判断系统大小端
func IsMachineLittleEndian() bool {
	var i int16 = 0x0001

	u := unsafe.Pointer(&i)
	b := *((*byte)(u))

	return b == 0x01
}

// IsWindows 当前操作系统是否Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux 当前操作系统是否Linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac 当前操作系统是否Mac OS/X.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}
