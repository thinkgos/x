package password

import (
	"errors"
)

const defaultPrivateKey = "@#$%^&*()"

// ErrCompareFailed compare failed
var ErrCompareFailed = errors.New("verify compare failed")

// Verify verify interface
type Verify interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}
