// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var (
	complexCharValues = map[string]string{
		"lower": `abcdefghijklmnopqrstuvwxyz`,
		"upper": `ABCDEFGHIJKLMNOPQRSTUVWXYZ`,
		"digit": `0123456789`,
		"spec":  ` !"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`",
	}

	complexChars        = complexCharValues["lower"] + complexCharValues["upper"] + complexCharValues["digit"]
	complexRequiredList []string
)

// SetupComplexity setup complexity with values
// value can set lower,upper,digit,spec
// default use lower, upper and digit to generate a random password, and not meets complexity.
func SetupComplexity(values []string) {
	complexChars = ""
	complexRequiredList = nil
	if len(values) != 1 || values[0] != "off" {
		for _, val := range values {
			if chars, ok := complexCharValues[val]; ok {
				complexChars += chars
				complexRequiredList = append(complexRequiredList, chars)
			}
		}
		if len(complexRequiredList) == 0 {
			// No valid character classes found; use all classes as default
			for _, chars := range complexCharValues {
				complexChars += chars
				complexRequiredList = append(complexRequiredList, chars)
			}
		}
	}

	if complexChars == "" {
		// No complexities to check; provide a sensible default for password generation
		complexChars = complexCharValues["lower"] + complexCharValues["upper"] + complexCharValues["digit"]
	}
}

// IsComplexEnough return True if password meets complexity settings
func IsComplexEnough(pwd string) bool {
	if len(complexChars) > 0 {
		for _, chars := range complexRequiredList {
			if !strings.ContainsAny(chars, pwd) {
				return false
			}
		}
	}
	return true
}

// Generate a random password
func Generate(n int) (string, error) {
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(complexChars)))
	for {
		for j := 0; j < n; j++ {
			rnd, err := rand.Int(rand.Reader, max)
			if err != nil {
				return "", err
			}
			buffer[j] = complexChars[rnd.Int64()]
		}
		if IsComplexEnough(string(buffer)) && string(buffer[0]) != " " && string(buffer[n-1]) != " " {
			return string(buffer), nil
		}
	}
}
