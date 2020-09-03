package extssh

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// ParsePrivateKey parse private key
func ParsePrivateKey(pemBytes []byte, passPhrase ...[]byte) (ssh.Signer, error) {
	if len(passPhrase) > 0 {
		return ssh.ParsePrivateKeyWithPassphrase(pemBytes, passPhrase[0])
	}
	return ssh.ParsePrivateKey(pemBytes)
}

// LoadPrivateKey load from file and parse private key
func LoadPrivateKey(filename string, passPhrase ...[]byte) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read rsa key file %w", err)
	}
	return ParsePrivateKey(key, passPhrase...)
}

// LoadPrivateKey2AuthMethod load from file and parse private key file to ssh.AuthMethod
func LoadPrivateKey2AuthMethod(filename string, passPhrase ...[]byte) (ssh.AuthMethod, error) {
	signer, err := LoadPrivateKey(filename, passPhrase...)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}
