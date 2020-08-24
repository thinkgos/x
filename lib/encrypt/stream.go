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
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"

	"gitlab.com/yawning/chacha20.git"
	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/salsa20/salsa"
)

// CipherInfo cipher information
type CipherInfo struct {
	KeyLen    int
	IvLen     int
	NewStream func(key, iv []byte, isEncode bool) (cipher.Stream, error)
}

var ciphers = map[string]CipherInfo{
	"aes-128-cfb":   {16, 16, newAESCFBStream},
	"aes-192-cfb":   {24, 16, newAESCFBStream},
	"aes-256-cfb":   {32, 16, newAESCFBStream},
	"aes-128-ctr":   {16, 16, newAESCTRStream},
	"aes-192-ctr":   {24, 16, newAESCTRStream},
	"aes-256-ctr":   {32, 16, newAESCTRStream},
	"des-cfb":       {8, 8, newDESStream},
	"bf-cfb":        {16, 8, newBlowFishStream},
	"cast5-cfb":     {16, 8, newCast5Stream},
	"rc4-md5":       {16, 16, newRC4MD5Stream},
	"rc4-md5-6":     {16, 6, newRC4MD5Stream},
	"chacha20":      {32, 8, newChaCha20Stream},
	"chacha20-ietf": {32, 12, newChaCha20IETFStream},
	"salsa20":       {32, 8, newSalsa20Stream},
}

// GetCipherInfo 根据方法获得 Cipher information
func GetCipherInfo(method string) (info CipherInfo, ok bool) {
	info, ok = ciphers[method]
	return
}

// CipherMethods 获取Cipher的所有支持方法
func CipherMethods() []string {
	keys := make([]string, 0, len(ciphers))
	for k := range ciphers {
		keys = append(keys, k)
	}
	return keys
}

// HasCipherMethod 是否有method方法
func HasCipherMethod(method string) (ok bool) {
	_, ok = ciphers[method]
	return
}

func newStream(block cipher.Block, iv []byte, isEncode bool) (cipher.Stream, error) {
	if isEncode {
		return cipher.NewCFBEncrypter(block, iv), nil
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

func newAESCFBStream(key, iv []byte, isEncode bool) (cipher.Stream, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newStream(block, iv, isEncode)
}

func newAESCTRStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCTR(block, iv), nil
}

func newDESStream(key, iv []byte, isEncode bool) (cipher.Stream, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newStream(block, iv, isEncode)
}

func newBlowFishStream(key, iv []byte, isEncode bool) (cipher.Stream, error) {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newStream(block, iv, isEncode)
}

func newCast5Stream(key, iv []byte, isEncode bool) (cipher.Stream, error) {
	block, err := cast5.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newStream(block, iv, isEncode)
}

func newRC4MD5Stream(key, iv []byte, _ bool) (cipher.Stream, error) {
	h := md5.New()
	h.Write(key) // nolint: errcheck
	h.Write(iv)  // nolint: errcheck
	return rc4.NewCipher(h.Sum(nil))
}

func newChaCha20Stream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return chacha20.New(key, iv)
}

func newChaCha20IETFStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return chacha20.New(key, iv)
}

func newSalsa20Stream(key, iv []byte, _ bool) (cipher.Stream, error) {
	var c salsaStreamCipher
	copy(c.nonce[:], iv[:8])
	copy(c.key[:], key[:32])
	return &c, nil
}

type salsaStreamCipher struct {
	nonce   [8]byte
	key     [32]byte
	counter int
}

func (c *salsaStreamCipher) XORKeyStream(dst, src []byte) {
	var buf []byte
	padLen := c.counter % 64
	dataSize := len(src) + padLen
	if cap(dst) >= dataSize {
		buf = dst[:dataSize]
	} else {
		buf = make([]byte, dataSize)
	}

	var subNonce [16]byte
	copy(subNonce[:], c.nonce[:])
	binary.LittleEndian.PutUint64(subNonce[len(c.nonce):], uint64(c.counter/64))

	// It's difficult to avoid data copy here. src or dst maybe slice from
	// Conn.Read/Write, which can't have padding.
	copy(buf[padLen:], src[:])
	salsa.XORKeyStream(buf, buf, &subNonce, &c.key)
	copy(dst, buf[padLen:])

	c.counter += len(src)
}
