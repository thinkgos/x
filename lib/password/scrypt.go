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
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/scrypt"
)

const maxSaltSize = 16

var _ Crypt = SCrypt{}

// SCrypt scrypt password encryption
type SCrypt struct{}

// GenerateFromPassword password hash encryption
func (SCrypt) GenerateFromPassword(password string) (string, error) {
	unencodedSalt := make([]byte, maxSaltSize)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		return "", err
	}

	rb, err := scrypt.Key([]byte(password), unencodedSalt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(append(unencodedSalt, rb...)), nil
}

// CompareHashAndPassword password hash verification
func (SCrypt) CompareHashAndPassword(hashedPassword, password string) error {
	orgRb, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return err
	}

	if len(orgRb) < maxSaltSize {
		return ErrCompareFailed
	}
	unencodedSalt := orgRb[:maxSaltSize]
	rb, err := scrypt.Key([]byte(password), unencodedSalt, 16384, 8, 1, 32)
	if err != nil {
		return err
	}

	if !bytes.Equal(orgRb[maxSaltSize:], rb) {
		return ErrCompareFailed
	}
	return nil
}
