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

// Package encrypt implement common encrypt and decrypt
package encrypt

import (
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"strconv"
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
// 		aes-128-ofb
// 		aes-192-ofb
// 		aes-256-ofb
// 		des-cfb
// 		des-ctr
// 		des-ofb
// 		blowfish-cfb
// 		blowfish-ctr
// 		blowfish-ofb
// 		cast5-cfb
// 		cast5-ctr
// 		cast5-ofb
// 		twofish-128-cfb
// 		twofish-192-cfb
// 		twofish-256-cfb
// 		twofish-128-ctr
// 		twofish-192-ctr
// 		twofish-256-ctr
// 		twofish-128-ofb
// 		twofish-192-ofb
// 		twofish-256-ofb
// 		tea-cfb
// 		tea-ctr
// 		tea-ofb
// 		xtea-cfb
// 		xtea-ctr
// 		xtea-ofb
// 		rc4-md5
// 		rc4-md5-6
// 		chacha20
// 		chacha20-ietf
// 		salsa20
func NewCipher(method, password string) (*Cipher, error) {
	if password == "" {
		return nil, errors.New("password required")
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

// NewStream new stream
func NewStream(method string, key, iv []byte, encrypt bool) (cipher.Stream, error) {
	info, ok := GetCipherInfo(method)
	if !ok {
		return nil, errors.New("Unsupported encryption method: " + method)
	}
	if len(key) < info.KeyLen {
		return nil, errors.New("invalid key size " + strconv.Itoa(len(key)))
	}
	if len(iv) < info.IvLen {
		return nil, errors.New("invalid IV length " + strconv.Itoa(len(iv)))
	}
	return info.NewStream(key[:info.KeyLen], iv[:info.IvLen], encrypt)
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
