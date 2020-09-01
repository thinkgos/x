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

type encDec struct {
	key       []byte
	iv        []byte
	newCipher func(key []byte) (cipher.Block, error)
	newStream func(block cipher.Block, iv []byte) cipher.Stream
}

// CipherInfo cipher information
type CipherInfo struct {
	keyLen        int
	ivLen         int
	newStream     func(*encDec) (cipher.Stream, error)
	newCipher     func(key []byte) (cipher.Block, error)
	newStreamFunc func(encrypt bool) func(block cipher.Block, iv []byte) cipher.Stream
}

var ciphers = map[string]CipherInfo{
	"aes-128-cfb": {
		16, 16, newStreamWithCipher,
		aes.NewCipher, newCfbStreamFunc,
	},
	"aes-192-cfb": {
		24, 16, newStreamWithCipher,
		aes.NewCipher, newCfbStreamFunc,
	},
	"aes-256-cfb": {
		32, 16, newStreamWithCipher,
		aes.NewCipher, newCfbStreamFunc,
	},
	"aes-128-ctr": {
		16, 16, newStreamWithCipher,
		aes.NewCipher, newCtrStreamFunc,
	},
	"aes-192-ctr": {
		24, 16, newStreamWithCipher,
		aes.NewCipher, newCtrStreamFunc,
	},
	"aes-256-ctr": {
		32, 16, newStreamWithCipher,
		aes.NewCipher, newCtrStreamFunc,
	},
	"aes-128-ofb": {
		16, 16, newStreamWithCipher,
		aes.NewCipher, newOfbStreamFunc,
	},
	"aes-192-ofb": {
		24, 16, newStreamWithCipher,
		aes.NewCipher, newOfbStreamFunc,
	},
	"aes-256-ofb": {
		32, 16, newStreamWithCipher,
		aes.NewCipher, newOfbStreamFunc,
	},
	"des-cfb": {
		8, 8, newStreamWithCipher,
		des.NewCipher, newCfbStreamFunc,
	},
	"des-ctr": {
		8, 8, newStreamWithCipher,
		des.NewCipher, newCtrStreamFunc,
	},
	"des-ofb": {
		8, 8, newStreamWithCipher,
		des.NewCipher, newOfbStreamFunc,
	},
	"3des-cfb": {
		24, 8, newStreamWithCipher,
		des.NewTripleDESCipher, newCfbStreamFunc,
	},
	"3des-ctr": {
		24, 8, newStreamWithCipher,
		des.NewTripleDESCipher, newCtrStreamFunc,
	},
	"3des-ofb": {
		24, 8, newStreamWithCipher,
		des.NewTripleDESCipher, newOfbStreamFunc,
	},
	"blowfish-cfb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) }, newCfbStreamFunc},
	"blowfish-ctr": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) }, newCtrStreamFunc},
	"blowfish-ofb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return blowfish.NewCipher(k) }, newOfbStreamFunc},
	"cast5-cfb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) }, newCfbStreamFunc,
	},
	"cast5-ctr": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) }, newCtrStreamFunc,
	},
	"cast5-ofb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return cast5.NewCipher(k) }, newOfbStreamFunc,
	},
	"twofish-128-cfb": {
		16, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCfbStreamFunc,
	},
	"twofish-192-cfb": {
		24, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCfbStreamFunc,
	},
	"twofish-256-cfb": {
		32, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCfbStreamFunc,
	},
	"twofish-128-ctr": {
		16, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCtrStreamFunc,
	},
	"twofish-192-ctr": {
		24, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCtrStreamFunc,
	},
	"twofish-256-ctr": {
		32, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newCtrStreamFunc,
	},
	"twofish-128-ofb": {
		16, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newOfbStreamFunc,
	},
	"twofish-192-ofb": {
		24, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newOfbStreamFunc,
	},
	"twofish-256-ofb": {
		32, 16, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return twofish.NewCipher(k) }, newOfbStreamFunc,
	},
	"tea-cfb": {
		16, 8, newStreamWithCipher,
		tea.NewCipher, newCfbStreamFunc,
	},
	"tea-ctr": {
		16, 8, newStreamWithCipher,
		tea.NewCipher, newCtrStreamFunc,
	},
	"tea-ofb": {
		16, 8, newStreamWithCipher,
		tea.NewCipher, newOfbStreamFunc,
	},
	"xtea-cfb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) }, newCfbStreamFunc,
	},
	"xtea-ctr": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) }, newCtrStreamFunc,
	},
	"xtea-ofb": {
		16, 8, newStreamWithCipher,
		func(k []byte) (cipher.Block, error) { return xtea.NewCipher(k) }, newOfbStreamFunc,
	},
	"rc4-md5": {
		16, 16, newRc4Md5Stream,
		nil, empty,
	},
	"rc4-md5-6": {
		16, 6, newRc4Md5Stream,
		nil, empty,
	},
	"chacha20": {
		32, 12, newChaCha20Stream,
		nil, empty,
	},
	"chacha20-ietf": {
		32, 24, newChaCha20IETFStream,
		nil, empty,
	},
	"salsa20": {
		32, 8, newSalsa20Stream,
		nil, empty,
	},
}

// GetCipherInfo 根据方法获得 Cipher information
func GetCipherInfo(method string) (info CipherInfo, ok bool) {
	info, ok = ciphers[method]
	return
}

// KeyLen return key len
func (sf *CipherInfo) KeyLen() int { return sf.keyLen }

// IvLen return iv len
func (sf *CipherInfo) IvLen() int { return sf.ivLen }

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

func newStreamWithCipher(ec *encDec) (cipher.Stream, error) {
	block, err := ec.newCipher(ec.key)
	if err != nil {
		return nil, err
	}
	return ec.newStream(block, ec.iv), nil
}

func newRc4Md5Stream(ec *encDec) (cipher.Stream, error) {
	h := md5.New()
	h.Write(ec.key) // nolint: errcheck
	h.Write(ec.iv)  // nolint: errcheck
	return rc4.NewCipher(h.Sum(nil))
}

func newChaCha20Stream(ec *encDec) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(ec.key, ec.iv)
}

func newChaCha20IETFStream(ec *encDec) (cipher.Stream, error) {
	return chacha20.NewUnauthenticatedCipher(ec.key, ec.iv)
}

func newSalsa20Stream(ec *encDec) (cipher.Stream, error) {
	var c salsaStreamCipher
	copy(c.nonce[:], ec.iv[:8])
	copy(c.key[:], ec.key[:32])
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

func newCfbStreamFunc(encrypt bool) func(block cipher.Block, iv []byte) cipher.Stream {
	if encrypt {
		return cipher.NewCFBEncrypter
	}
	return cipher.NewCFBDecrypter
}

func newCtrStreamFunc(bool) func(block cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewCTR
}

func newOfbStreamFunc(bool) func(block cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewOFB
}

func empty(bool) func(block cipher.Block, iv []byte) cipher.Stream {
	return func(block cipher.Block, iv []byte) cipher.Stream { return nil }
}
