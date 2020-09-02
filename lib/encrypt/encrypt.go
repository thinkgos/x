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

// Package encrypt implement common encrypt and decrypt for stream
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"strconv"

	"golang.org/x/crypto/tea"
)

// KeyIvLen key and iv length interface
type KeyIvLen interface {
	KeyLen() int
	IvLen() int
}

// complexCipher cipher information
type complexCipher struct {
	keyLen     int
	ivLen      int
	newCipher  func(key []byte) (cipher.Block, error)
	newEncrypt func(block cipher.Block, iv []byte) cipher.Stream
	newDecrypt func(block cipher.Block, iv []byte) cipher.Stream
}

// KeyLen return key len
func (sf complexCipher) KeyLen() int { return sf.keyLen }

// IvLen return iv len
func (sf complexCipher) IvLen() int { return sf.ivLen }

// simpleCiphers cipher information
type simpleCipher struct {
	keyLen    int
	ivLen     int
	newStream func(key, iv []byte) (cipher.Stream, error)
}

// KeyLen return key len
func (sf simpleCipher) KeyLen() int { return sf.keyLen }

// IvLen return iv len
func (sf simpleCipher) IvLen() int { return sf.ivLen }

var complexCiphers = map[string]complexCipher{
	"aes-128-cfb":     {16, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"aes-192-cfb":     {24, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"aes-256-cfb":     {32, 16, aes.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"aes-128-ctr":     {16, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR},
	"aes-192-ctr":     {24, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR},
	"aes-256-ctr":     {32, 16, aes.NewCipher, cipher.NewCTR, cipher.NewCTR},
	"aes-128-ofb":     {16, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB},
	"aes-192-ofb":     {24, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB},
	"aes-256-ofb":     {32, 16, aes.NewCipher, cipher.NewOFB, cipher.NewOFB},
	"des-cfb":         {8, 8, des.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"des-ctr":         {8, 8, des.NewCipher, cipher.NewCTR, cipher.NewCTR},
	"des-ofb":         {8, 8, des.NewCipher, cipher.NewOFB, cipher.NewOFB},
	"3des-cfb":        {24, 8, des.NewTripleDESCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"3des-ctr":        {24, 8, des.NewTripleDESCipher, cipher.NewCTR, cipher.NewCTR},
	"3des-ofb":        {24, 8, des.NewTripleDESCipher, cipher.NewOFB, cipher.NewOFB},
	"blowfish-cfb":    {16, 8, NewBlowfishCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"blowfish-ctr":    {16, 8, NewBlowfishCipher, cipher.NewCTR, cipher.NewCTR},
	"blowfish-ofb":    {16, 8, NewBlowfishCipher, cipher.NewOFB, cipher.NewOFB},
	"cast5-cfb":       {16, 8, NewCast5Cipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"cast5-ctr":       {16, 8, NewCast5Cipher, cipher.NewCTR, cipher.NewCTR},
	"cast5-ofb":       {16, 8, NewCast5Cipher, cipher.NewOFB, cipher.NewOFB},
	"twofish-128-cfb": {16, 16, NewTwofishCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"twofish-192-cfb": {24, 16, NewTwofishCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"twofish-256-cfb": {32, 16, NewTwofishCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"twofish-128-ctr": {16, 16, NewTwofishCipher, cipher.NewCTR, cipher.NewCTR},
	"twofish-192-ctr": {24, 16, NewTwofishCipher, cipher.NewCTR, cipher.NewCTR},
	"twofish-256-ctr": {32, 16, NewTwofishCipher, cipher.NewCTR, cipher.NewCTR},
	"twofish-128-ofb": {16, 16, NewTwofishCipher, cipher.NewOFB, cipher.NewOFB},
	"twofish-192-ofb": {24, 16, NewTwofishCipher, cipher.NewOFB, cipher.NewOFB},
	"twofish-256-ofb": {32, 16, NewTwofishCipher, cipher.NewOFB, cipher.NewOFB},
	"xtea-cfb":        {16, 8, NewXteaCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"xtea-ctr":        {16, 8, NewXteaCipher, cipher.NewCTR, cipher.NewCTR},
	"xtea-ofb":        {16, 8, NewXteaCipher, cipher.NewOFB, cipher.NewOFB},
	"tea-cfb":         {16, 8, tea.NewCipher, cipher.NewCFBEncrypter, cipher.NewCFBDecrypter},
	"tea-ctr":         {16, 8, tea.NewCipher, cipher.NewCTR, cipher.NewCTR},
	"tea-ofb":         {16, 8, tea.NewCipher, cipher.NewOFB, cipher.NewOFB},
}

var simpleCiphers = map[string]simpleCipher{
	"rc4-md5":       {16, 16, NewRc4Md5},
	"rc4-md5-6":     {16, 6, NewRc4Md5},
	"chacha20":      {32, 12, NewChacha20},
	"chacha20-ietf": {32, 24, NewChacha20},
	"salsa20":       {32, 8, NewSalsa20},
}

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
// 		3des-cfb
// 		3des-ctr
// 		3des-ofb
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

	if info, ok := complexCiphers[method]; ok {
		key := Evp2Key(password, info.keyLen)

		// hash(key) -> read IV
		riv := sha256.New().Sum(key)[:info.IvLen()]
		rd, err := (&Stream{info.newEncrypt}).New(key, riv, info.newCipher)
		if err != nil {
			return nil, err
		}
		// hash(read IV) -> write IV
		wiv := sha256.New().Sum(riv)[:info.IvLen()]
		wr, err := (&Stream{info.newDecrypt}).New(key, wiv, info.newCipher)
		if err != nil {
			return nil, err
		}
		return &Cipher{wr, rd}, nil
	}

	if info, ok := simpleCiphers[method]; ok {
		key := Evp2Key(password, info.keyLen)

		// hash(key) -> read IV
		riv := sha256.New().Sum(key)[:info.IvLen()]
		wr, err := info.newStream(key[:info.keyLen], riv[:info.ivLen])
		if err != nil {
			return nil, err
		}
		// hash(read IV) -> write IV
		wiv := sha256.New().Sum(riv)[:info.IvLen()]
		rd, err := info.newStream(key[:info.keyLen], wiv[:info.ivLen])
		if err != nil {
			return nil, err
		}
		return &Cipher{wr, rd}, nil
	}
	return nil, errors.New("unsupported encryption method: " + method)
}

// NewStream new stream
func NewStream(method string, key, iv []byte, encrypt bool) (cipher.Stream, error) {
	check := func(info KeyIvLen) error {
		if len(key) < info.KeyLen() {
			return errors.New("invalid key size " + strconv.Itoa(len(key)))
		}
		if len(iv) < info.IvLen() {
			return errors.New("invalid IV length " + strconv.Itoa(len(iv)))
		}
		return nil
	}

	if info, ok := complexCiphers[method]; ok {
		if err := check(info); err != nil {
			return nil, err
		}
		encdec := info.newDecrypt
		if encrypt {
			encdec = info.newEncrypt
		}
		return (&Stream{encdec}).New(key[:info.keyLen], iv[:info.ivLen], info.newCipher)
	}

	if info, ok := simpleCiphers[method]; ok {
		if err := check(info); err != nil {
			return nil, err
		}
		return info.newStream(key[:info.keyLen], iv[:info.ivLen])
	}
	return nil, errors.New("unsupported encryption method: " + method)
}

// GetCipher 根据方法获得 Cipher information
func GetCipher(method string) (KeyIvLen, bool) {
	if info, ok := complexCiphers[method]; ok {
		return info, ok
	}
	info, ok := simpleCiphers[method]
	return info, ok
}

// CipherMethods 获取Cipher的所有支持方法
func CipherMethods() []string {
	keys := make([]string, 0, len(complexCiphers)+len(simpleCiphers))
	for k := range complexCiphers {
		keys = append(keys, k)
	}
	for k := range simpleCiphers {
		keys = append(keys, k)
	}
	return keys
}

// HasCipherMethod 是否有method方法
func HasCipherMethod(method string) (ok bool) {
	if _, ok = complexCiphers[method]; !ok {
		_, ok = simpleCiphers[method]
	}
	return
}

// Valid method password is valid or not
func Valid(method, password string) bool {
	_, err := NewCipher(method, password)
	return err == nil
}

// Evp2Key evp to key
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
