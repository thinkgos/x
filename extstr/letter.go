package extstr

import (
	"unicode"
)

// IsLetter 是否是英文字母
func IsLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// IsLetters 字符串是否全英文字母
func IsLetters(str string) bool {
	for _, r := range str {
		if !IsLetter(r) {
			return false
		}
	}
	return str != ""
}

// HasLetter 字符串是否含有(英文)字母.
func HasLetter(str string) bool {
	for _, r := range str {
		if IsLetter(r) {
			return true
		}
	}
	return false
}

// IsASCII 是否IsASCII字符串.
func IsASCII(str string) bool {
	for _, r := range str {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return str != ""
}

// HasChinese 字符串是否含有中文.
func HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func HasSpecialChar(str string) bool {
	for _, r := range str {
		// IsPunct 判断 r 是否为一个标点字符 (类别 P)
		// IsSymbol 判断 r 是否为一个符号字符
		// IsMark 判断 r 是否为一个 mark 字符 (类别 M)
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			return true
		}
	}
	return false
}
