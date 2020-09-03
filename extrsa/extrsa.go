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

package extrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

/**
生成私钥流程:
	1. 使用rsa中的GenerateKey方法生成私钥(bits建议1024的整数倍):
		func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)
	2. 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
		func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	3. 将私钥字符串设置到pem格式块中,初始化一个pem.Block块
		type Block struct {
		   Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
		   Headers map[string]string // 可选的头项
		   Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
		}
	4. 通过pem将设置好的数据进行编码
		func Encode(out io.Writer, b *Block) error

生成公钥流程:
	1. 从得到的私钥对象中将公钥提出
		PrivateKey.PublicKey
	2. 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
		func MarshalPKIXPublicKey(pub interface{}) ([]byte, error)
	3. 将公钥字符串设置到pem格式块中,初始化一个pem.Block块
		type Block struct {
		   Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
		   Headers map[string]string // 可选的头项
		   Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
		}
	4. 通过pem将设置好的数据进行编码
		func Encode(out io.Writer, b *Block) error

加解密:
	加密: func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
	解密: func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)

数字签名加验签:
	加签: func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
	验签: func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
*/

// GenerateKeys generate keys return private and public pem.Block.
// bits: 1024*n, n > 0
func GenerateKeys(bits int) (private, public *pem.Block, err error) {
	// generate rand privateKey
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	// 通过x509将ras私钥序列化为ASN.1 的 DER编码字符串
	x50PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	// 通过x509将ras公钥序列化为ASN.1 的 DER编码字符串
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x50PrivateKey},
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509PublicKey}, nil
}

// ParsePublicKey parse public key
func ParsePublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("invalid input public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := pub.(*rsa.PublicKey)
	return publicKey, nil
}

// LoadPublicKey load public key
func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParsePublicKey(b)
}

// ParsePrivateKey parse  private key
func ParsePrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("invalid input private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadPrivateKey load private key
func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKey(b)
}

// WriteFile write file
func WriteFile(filename string, block *pem.Block) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return pem.Encode(f, block)
}
