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

	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/salsa20/salsa"
	"golang.org/x/crypto/tea"
	"golang.org/x/crypto/twofish"
	"golang.org/x/crypto/xtea"

	"golang.org/x/crypto/chacha20"
)

// CipherInfo cipher information
type CipherInfo struct {
	KeyLen    int
	IvLen     int
	NewStream func(key, iv []byte, encrypt bool) (cipher.Stream, error)
}

var ciphers = map[string]CipherInfo{
	"aes-128-cfb":     {16, 16, newAesCfbStream},
	"aes-192-cfb":     {24, 16, newAesCfbStream},
	"aes-256-cfb":     {32, 16, newAesCfbStream},
	"aes-128-ctr":     {16, 16, newAesCtrStream},
	"aes-192-ctr":     {24, 16, newAesCtrStream},
	"aes-256-ctr":     {32, 16, newAesCtrStream},
	"aes-128-ofb":     {16, 16, newAesOfbStream},
	"aes-192-ofb":     {24, 16, newAesOfbStream},
	"aes-256-ofb":     {32, 16, newAesOfbStream},
	"des-cfb":         {8, 8, newDesCfbStream},
	"des-ctr":         {8, 8, newDesCtrStream},
	"des-ofb":         {8, 8, newDesOfbStream},
	"blowfish-cfb":    {16, 8, newBlowfishCfbStream},
	"blowfish-ctr":    {16, 8, newBlowfishCtrStream},
	"blowfish-ofb":    {16, 8, newBlowfishOfbStream},
	"cast5-cfb":       {16, 8, newCast5CfbStream},
	"cast5-ctr":       {16, 8, newCast5CtrStream},
	"cast5-ofb":       {16, 8, newCast5OfbStream},
	"twofish-128-cfb": {16, 16, newTwofishCfbStream},
	"twofish-192-cfb": {24, 16, newTwofishCfbStream},
	"twofish-256-cfb": {32, 16, newTwofishCfbStream},
	"twofish-128-ctr": {16, 16, newTwofishCtrStream},
	"twofish-192-ctr": {24, 16, newTwofishCtrStream},
	"twofish-256-ctr": {32, 16, newTwofishCtrStream},
	"twofish-128-ofb": {16, 16, newTwofishOfbStream},
	"twofish-192-ofb": {24, 16, newTwofishOfbStream},
	"twofish-256-ofb": {32, 16, newTwofishOfbStream},
	"tea-cfb":         {16, 8, newTeaCfbStream},
	"tea-ctr":         {16, 8, newTeaCtrStream},
	"tea-ofb":         {16, 8, newTeaOfbStream},
	"xtea-cfb":        {16, 8, newXteaCfbStream},
	"xtea-ctr":        {16, 8, newXteaCtrStream},
	"xtea-ofb":        {16, 8, newXteaOfbStream},
	"rc4-md5":         {16, 16, newRc4Md5Stream},
	"rc4-md5-6":       {16, 6, newRc4Md5Stream},
	"chacha20":        {32, 12, newChaCha20Stream},
	"chacha20-ietf":   {32, 24, newChaCha20IETFStream},
	"salsa20":         {32, 8, newSalsa20Stream},
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

func newCfbStream(newCipher func(k []byte) (cipher.Block, error), key, iv []byte, encrypt bool) (cipher.Stream, error) {
	block, err := newCipher(key)
	if err != nil {
		return nil, err
	}
	if encrypt {
		return cipher.NewCFBEncrypter(block, iv), nil
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

func newCtrStream(newCipher func(k []byte) (cipher.Block, error), key, iv []byte) (cipher.Stream, error) {
	block, err := newCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCTR(block, iv), nil
}

func newOfbStream(newCipher func(k []byte) (cipher.Block, error), key, iv []byte) (cipher.Stream, error) {
	block, err := newCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewOFB(block, iv), nil
}

func newAesCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(aes.NewCipher, key, iv, encrypt)
}

func newAesCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(aes.NewCipher, key, iv)
}

func newAesOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(aes.NewCipher, key, iv)
}

func newDesCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(des.NewCipher, key, iv, encrypt)
}

func newDesCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(des.NewCipher, key, iv)
}

func newDesOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(des.NewCipher, key, iv)
}

func newBlowfishCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(func(k []byte) (cipher.Block, error) {
		return blowfish.NewCipher(k)
	}, key, iv, encrypt)
}

func newBlowfishCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(func(k []byte) (cipher.Block, error) {
		return blowfish.NewCipher(k)
	}, key, iv)
}

func newBlowfishOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(func(k []byte) (cipher.Block, error) {
		return blowfish.NewCipher(k)
	}, key, iv)
}

func newCast5CfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(func(k []byte) (cipher.Block, error) {
		return cast5.NewCipher(k)
	}, key, iv, encrypt)
}

func newCast5CtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(func(k []byte) (cipher.Block, error) {
		return cast5.NewCipher(k)
	}, key, iv)
}

func newCast5OfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(func(k []byte) (cipher.Block, error) {
		return cast5.NewCipher(k)
	}, key, iv)
}
func newTwofishCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(func(k []byte) (cipher.Block, error) {
		return twofish.NewCipher(k)
	}, key, iv, encrypt)
}

func newTwofishCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(func(k []byte) (cipher.Block, error) {
		return twofish.NewCipher(k)
	}, key, iv)
}
func newTwofishOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(func(k []byte) (cipher.Block, error) {
		return twofish.NewCipher(k)
	}, key, iv)
}

func newTeaCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(tea.NewCipher, key, iv, encrypt)
}

func newTeaCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(tea.NewCipher, key, iv)
}
func newTeaOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(tea.NewCipher, key, iv)
}

func newXteaCfbStream(key, iv []byte, encrypt bool) (cipher.Stream, error) {
	return newCfbStream(func(k []byte) (cipher.Block, error) {
		return xtea.NewCipher(k)
	}, key, iv, encrypt)
}

func newXteaCtrStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newCtrStream(func(k []byte) (cipher.Block, error) {
		return xtea.NewCipher(k)
	}, key, iv)
}

func newXteaOfbStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return newOfbStream(func(k []byte) (cipher.Block, error) {
		return xtea.NewCipher(k)
	}, key, iv)
}

func newRc4Md5Stream(key, iv []byte, _ bool) (cipher.Stream, error) {
	h := md5.New()
	h.Write(key) // nolint: errcheck
	h.Write(iv)  // nolint: errcheck
	return rc4.NewCipher(h.Sum(nil))
}

func newChaCha20Stream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(key, iv)
}

func newChaCha20IETFStream(key, iv []byte, _ bool) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(key, iv)
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
	copy(buf[padLen:], src)
	salsa.XORKeyStream(buf, buf, &subNonce, &c.key)
	copy(dst, buf[padLen:])

	c.counter += len(src)
}
