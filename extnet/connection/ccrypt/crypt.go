// Package ccrypt 实现net.conn的加密conn接口 aes cfb加密码的连接,通过提供的配置使用pbkdf2生成key,
// 依靠key和hash生成iv
package ccrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"hash"
	"io"
	"net"

	"golang.org/x/crypto/pbkdf2"
)

// Config defaults
const (
	DefaultIterations = 2048
	DefaultKeySize    = 32 // 32byte 256bits
)

// Config defaults
var (
	DefaultHash = sha256.New
	DefaultSalt = []byte(`dfaod..ga'.';dfa,;dfa',';023pkadb;j90[[q93jr4poemf';adm;lcv;aodfm'a;dlm[f09eqegoain;lfgankd;gihqp8e8hladknaldkpq3h4''`)
)

// Config PBKDF2 key 生成的参数配置
type Config struct {
	Password   string           // 用户加密密钥
	Salt       []byte           // default: DefaultSalt
	Iterations int              // 迭代次数,次数越多加密与解密所需时间越长, default: 2048
	KeySize    int              // 期望密文长度,支持16,24,32, default: 32
	Hash       func() hash.Hash // 加密所使用的hash函数, default: DefaultHash(sha256.New)
}

// Conn conn with ...
type Conn struct {
	net.Conn
	r io.Reader
	w io.Writer
}

//New 创建一个aes cfb加密码的连接,通过提供的配置使用pbkdf2生成key,依靠key和hash生成iv
func New(c net.Conn, cfg Config) *Conn {
	//set defaults
	if len(cfg.Salt) == 0 {
		cfg.Salt = DefaultSalt
	}
	if cfg.Iterations == 0 {
		cfg.Iterations = DefaultIterations
	}
	if cfg.KeySize != 16 && cfg.KeySize != 24 && cfg.KeySize != 32 {
		cfg.KeySize = DefaultKeySize
	}
	if cfg.Hash == nil {
		cfg.Hash = DefaultHash
	}

	// generate key
	key := pbkdf2.Key([]byte(cfg.Password), cfg.Salt, cfg.Iterations, cfg.KeySize, cfg.Hash)

	//key will be always be the correct size so this will never error
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// hash(key) -> read IV
	riv := DefaultHash().Sum(key)
	// hash(read IV) -> write IV
	wiv := DefaultHash().Sum(riv)
	return &Conn{
		c,
		&cipher.StreamReader{
			S: cipher.NewCFBDecrypter(block, riv[:block.BlockSize()]),
			R: c,
		},
		&cipher.StreamWriter{
			S: cipher.NewCFBEncrypter(block, wiv[:block.BlockSize()]),
			W: c,
		},
	}
}

// Read reads data from the connection.
func (sf *Conn) Read(p []byte) (int, error) {
	return sf.r.Read(p)
}

// Write writes data to the connection.
func (sf *Conn) Write(p []byte) (int, error) {
	return sf.w.Write(p)
}
