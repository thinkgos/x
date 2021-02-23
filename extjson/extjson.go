package extjson

import (
	"encoding/json"

	"github.com/thinkgos/x/internal/bytesconv"
)

// MarshalToString convenient method to write as string instead of []byte
func MarshalToString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return bytesconv.Bytes2Str(b), nil
}

// MarshalIndentToString is like Marshal but applies Indent to format the output.
func MarshalIndentToString(v interface{}, prefix, indent string) (string, error) {
	b, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}
	return bytesconv.Bytes2Str(b), nil
}

// UnmarshalFromString is a convenient method to read from string instead of []byte
func UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
