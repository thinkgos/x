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
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

var _ Verify = SCrypt{}

type SCrypt struct{}

// Hash password hash encryption
func (sf SCrypt) Hash(password, salt string) (string, error) {
	rb, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(rb), nil
}

// Compare password hash verification
func (sf SCrypt) Compare(password, salt, hash string) error {
	h, err := sf.Hash(password, salt)
	if err != nil {
		return err
	}
	if hash != h {
		return ErrCompareFailed
	}
	return nil
}
