package extstr

import (
	"testing"

	"github.com/thinkgos/go-core-package/extrand"
)

func TestIsLetter(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			"a",
			'a',
			true,
		},
		{
			"z",
			'z',
			true,
		},
		{
			"A",
			'A',
			true,
		},
		{
			"Z",
			'Z',
			true,
		},
		{
			"no letter",
			'.',
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetter(tt.r); got != tt.want {
				t.Errorf("IsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLetters(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"空字符",
			"",
			false,
		},
		{
			"alpha",
			extrand.RandLetter(8),
			true,
		},
		{
			"chinese",
			"中国",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLetters(tt.str); got != tt.want {
				t.Errorf("IsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasLetter(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"空字符",
			"",
			false,
		},
		{
			"has",
			"Heelwo",
			true,
		},
		{
			"no has",
			"中国",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasLetter(tt.str); got != tt.want {
				t.Errorf("HasLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"", "", false},
		{"", "ｆｏｏbar", false},
		{"", "ｘｙｚ０９８", false},
		{"", "１２３456", false},
		{"", "你好，世界", false},
		{"", "foobar", true},
		{"", "0987654321", true},
		{"", "test@example.com", true},
		{"", "1234abcDEF", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsASCII(tt.str); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasChinese(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"空字符",
			"",
			false,
		},
		{
			"无汉字",
			"124.abc",
			false,
		},
		{
			"有汉字",
			"hello你好",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasChinese(tt.str); got != tt.want {
				t.Errorf("HasChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSpecialChar(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"",
			"",
			false,
		},
		{
			"",
			"`~!@#$%^&*()_+-=:'|<>?,./\"",
			true,
		},
		{
			"",
			"Hello ៉៊់៌៍！",
			true,
		},
		{
			"",
			"hello world",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSpecialChar(tt.str); got != tt.want {
				t.Errorf("HasSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
