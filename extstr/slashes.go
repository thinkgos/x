package extstr

import (
	"bytes"
)

// AddSlashes returns a string with backslashes added before characters
// that need to be escaped.
// 使用反斜线引用字符串,对' " \转义.
func AddSlashes(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		switch ch {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

// StripSlashes returns a string with backslashes stripped off. (\' becomes ' and so on.)
// Double backslashes (\\) are made into a single backslash (\).
// 使用反斜线反引用字符串,对 \' \" \\反转义.
func StripSlashes(s string) string {
	var buf bytes.Buffer

	l, skip := len(s), false
	for i, ch := range s {
		if skip {
			buf.WriteRune(ch)
			skip = false
			continue
		}

		if ch == '\\' {
			if i+1 < l && s[i+1] == '\\' {
				skip = true
			}
			continue
		}

		buf.WriteRune(ch)
	}
	return buf.String()
}

// QuoteMeta returns a version of str with a backslash character (\)
// before every character that is among these: . \ + * ? [ ^ ] ( $ )
// 转义元字符集,包括 . + \ ( $ ) [ ^ ] * ?
func QuoteMeta(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		switch ch {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}
