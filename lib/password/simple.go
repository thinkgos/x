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

package password

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

var _ Verify = Simple{}

// Simple simple password encryption
type Simple struct {
	key string
}

// NewSimple new simple password encryption with key
func NewSimple(key string) *Simple {
	if key == "" {
		key = defaultPrivateKey
	}
	return &Simple{key}
}

// Hash password hash encryption 加盐法 md5Pwd+`@#$%`+md5Pwd+`^&*()`拼接
func (sf Simple) Hash(password string) (string, error) {
	h := md5.New()
	_, _ = io.WriteString(h, password+sf.key)

	md5Pwd := hex.EncodeToString(h.Sum(nil))
	// 加盐值加密
	_, _ = io.WriteString(h, sf.key)
	_, _ = io.WriteString(h, password)
	_, _ = io.WriteString(h, `@#$%`+sf.key)
	_, _ = io.WriteString(h, md5Pwd)
	_, _ = io.WriteString(h, `^&*()`+sf.key)
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Compare password hash verification
func (sf Simple) Compare(password, hash string) error {
	h, err := sf.Hash(password)
	if err != nil {
		return err
	}
	if hash != h {
		return ErrCompareFailed
	}
	return nil
}
