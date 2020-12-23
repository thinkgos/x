package univ

import (
	"sort"
	"strings"
)

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (v Values) Set(key, value string) {
	v[key] = []string{value}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Del deletes the values associated with key.
func (v Values) Del(key string) {
	delete(v, key)
}

// ParseValue parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
func ParseValue(values, separator, delimiter string) Values {
	m := make(Values)
	for values != "" {
		key := values
		if i := strings.IndexAny(key, delimiter); i >= 0 {
			key, values = key[:i], key[i+1:]
		} else {
			values = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, separator); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		if key == "" || values == "" {
			continue
		}
		m[key] = append(m[key], value)
	}
	return m
}

// Encode encodes the values into form (separator="=" delimiter="&")
// ("bar=baz&foo=quux") sorted by key.
func (v Values) Encode(separator, delimiter string) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteString(delimiter)
			}
			buf.WriteString(k)
			buf.WriteString(separator)
			buf.WriteString(v)
		}
	}
	return buf.String()
}
