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

package encrypt

import (
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"errors"
)

// Cipher implement write and read cipher.Stream
type Cipher struct {
	Write cipher.Stream
	Read  cipher.Stream
}

// NewCipher new cipher
// method support:
// 		aes-128-cfb
// 		aes-192-cfb
// 		aes-256-cfb
// 		aes-128-ctr
// 		aes-192-ctr
// 		aes-256-ctr
// 		des-cfb
// 		bf-cfb
// 		cast5-cfb
// 		rc4-md5
// 		rc4-md5-6
// 		chacha20
// 		chacha20-ietf
func NewCipher(method, password string) (*Cipher, error) {
	if password == "" {
		return nil, errors.New("empty password")
	}
	info, ok := GetCipherInfo(method)
	if !ok {
		return nil, errors.New("Unsupported encryption method: " + method)
	}
	key := Evp2Key(password, info.KeyLen)

	//hash(key) -> read IV
	riv := sha256.New().Sum(key)[:info.IvLen]
	rd, err := info.NewStream(key, riv, false)
	if err != nil {
		return nil, err
	}
	//hash(read IV) -> write IV
	wiv := sha256.New().Sum(riv)[:info.IvLen]
	wr, err := info.NewStream(key, wiv, true)
	if err != nil {
		return nil, err
	}
	return &Cipher{wr, rd}, nil
}

// Evp2Key ...
func Evp2Key(password string, keyLen int) (key []byte) {
	const md5Len = 16

	cnt := (keyLen-1)/md5Len + 1
	m := make([]byte, cnt*md5Len)
	copy(m, md5sum([]byte(password)))

	// Repeatedly call md5 until bytes generated is enough.
	// Each call to md5 uses data: prev md5 sum + password.
	d := make([]byte, md5Len+len(password))
	for start, i := 0, 1; i < cnt; i++ {
		start += md5Len
		copy(d, m[start-md5Len:start])
		copy(d[md5Len:], password)
		copy(m[start:], md5sum(d))
	}
	return m[:keyLen]
}

func md5sum(b []byte) []byte {
	h := md5.New()
	h.Write(b) // nolint: errcheck
	return h.Sum(nil)
}
