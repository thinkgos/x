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
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

var _ Crypt = Simple{}

// Simple simple password encryption
type Simple struct{}

// Hash password hash encryption 加盐法
func (sf Simple) GenerateFromPassword(password string) (string, error) {
	unencodedSalt := make([]byte, maxSaltSize)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		return "", err
	}

	pwd := sf.hash(password, string(unencodedSalt))
	return base64.StdEncoding.EncodeToString(append(unencodedSalt, pwd...)), nil
}

// Compare password hash verification
func (sf Simple) CompareHashAndPassword(hashedPassword, password string) error {
	orgRb, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return err
	}

	if len(orgRb) < maxSaltSize {
		return ErrCompareFailed
	}
	unencodedSalt := orgRb[:maxSaltSize]

	pwd := sf.hash(password, string(unencodedSalt))

	if !bytes.Equal(orgRb[maxSaltSize:], pwd) {
		return ErrCompareFailed
	}
	return nil
}

const (
	salt1 = `@#$%`
	salt2 = `^&*()`
)

func (Simple) hash(password, salt string) []byte {
	md5Pwd := md5.Sum([]byte(password))

	build := new(bytes.Buffer)
	build.Grow(len(password) + len(salt)*3 + len(md5Pwd) + len(salt1) + len(salt2))
	// 加盐值加密
	build.WriteString(salt)
	build.WriteString(password)
	build.WriteString(salt1)
	build.WriteString(salt)
	build.Write(md5Pwd[:])
	build.WriteString(salt2)
	build.WriteString(salt)

	src := md5.Sum(build.Bytes())

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src[:])
	return dst
}
