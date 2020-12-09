// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package regex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmail(t *testing.T) {
	emails := map[string]bool{
		`test@example.com`:             true,
		`single-character@b.org`:       true,
		`uncommon_address@test.museum`: true,
		`local@sld.UPPER`:              true,
		`@missing.org`:                 false,
		`missing@.com`:                 false,
		`missing@qq.`:                  false,
		`wrong-ip@127.1.1.1.26`:        false,
	}
	for e, r := range emails {
		require.Equal(t, r, IsEmail(e))
	}
}

func TestIsEmailRFC(t *testing.T) {
	require.True(t, IsEmailRFC("test@example.com"))
}

func TestIsUrl(t *testing.T) {
	var tests = []struct {
		name string
		url  string
		want bool
	}{
		{"", "", false},
		{"", "http://foo.bar/#com", true},
		{"", "http://foobar.com", true},
		{"", "https://foobar.com", true},
		{"", "foobar.com", false},
		{"", "http://foobar.coffee/", true},
		{"", "http://foobar.中文网/", true},
		{"", "https://foobar.org/", true},
		{"", "http://foobar.org:8080/", true},
		{"", "ftp://foobar.ru/", true},
		{"", "http://user:pass@www.foobar.com/", true},
		{"", "http://127.0.0.1/", true},
		{"", "http://duckduckgo.com/?q=%2F", true},
		{"", "http://localhost:3000/", true},
		{"", "http://foobar.com/?foo=bar#baz=qux", true},
		{"", "http://foobar.com?foo=bar", true},
		{"", "http://www.xn--froschgrn-x9a.net/", true},
		{"", "invalid.", false},
		{"", ".com", false},
		{"", "http://www.foo_bar.com/", true},
		{"", "http://localhost:3000/", true},
		{"", "http://foobar.com/#baz=qux", true},
		{"", "http://foobar.com/t$-_.+!*\\'(),", true},
		{"", "http://www.foobar.com/~foobar", true},
		{"", "http://www.-foobar.com/", true},
		{"", "http://www.foo---bar.com/", true},
		{"", "/abs/test/dir", false},
		{"", "./rel/test/dir", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, IsURL(tt.url))
		})
	}
}

func BenchmarkIsEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("test@example.com")
	}
}

func BenchmarkIsUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("http://example.com")
	}
}

func TestIsRGBColor(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"", "", false},
		{"", "rgb(0,31,255)", true},
		{"", "rgb(1,349,275)", false},
		{"", "rgb(01,31,255)", false},
		{"", "rgb(0.6,31,255)", false},
		{"", "rgba(0,31,255)", false},
		{"", "rgb(0,  31, 255)", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRGBColor(tt.s); got != tt.want {
				t.Errorf("IsRGBColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  string
		want1 bool
	}{
		{"", "", "", false},
		{"", "#ff", "#ff", false},
		{"", "fff0", "fff0", false},
		{"", "#ff12FG", "#ff12FG", false},
		{"", "CCccCC", "#CCCCCC", true},
		{"", "fff", "#FFF", true},
		{"", "#f00", "#F00", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsHexColor(tt.s)
			if got != tt.want {
				t.Errorf("IsHexColor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsHexColor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"", "", false},
		{"", "12345", true},
		{"", "helloworld", true},
		{"", "PI314159", true},
		{"", "你好，世界", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumeric(tt.s); got != tt.want {
				t.Errorf("IsAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsWord(t *testing.T) {
	var tests = []struct {
		name string
		str  string
		want bool
	}{
		{"", "", false},
		{"", "_Football", false},
		{"", "-Football", false},
		{"", " 3.124", false},
		{"", "hello world.你好，世界！", false},
		{"", "世界", true},
		{"", "hello", true},
		{"", "作品T", true},
		{"", "8point", true},
		{"", "hello_Kitty2", true},
		{"", "hello-Kitty2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWord(tt.str); got != tt.want {
				t.Errorf("IsWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChinese(t *testing.T) {
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
			"全中文",
			"你好世界",
			true,
		},
		{
			"非中文",
			"你好hello",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.str); got != tt.want {
				t.Errorf("IsChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseName(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"", "", false},
		{"", "hello world", false},
		{"", "赵武", true},
		{"", "赵武a", false},
		{"", "南宫先生", true},
		{"", "吉乃•阿衣·依扎嫫", true},
		{"", "古丽莎•卡迪尔", true},
		{"", "迪丽热巴.迪力木拉提", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseName(tt.str); got != tt.want {
				t.Errorf("IsChineseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWhitespaces(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"", "abacaba", false},
		{"", "", false},
		{"", "    ", true},
		{"", "  \r\n  ", true},
		{"", "\014\012\011\013\015", true},
		{"", "\014\012\011\013 abc  \015", false},
		{"", "\f\n\t\v\r\f", true},
		{"", "x\n\t\t\t\t", false},
		{"", "\f\n\t  \n\n\n   \v\r\f", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWhitespaces(tt.str); got != tt.want {
				t.Errorf("IsWhitespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasWhitespace(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"", "abacaba", false},
		{"", "", false},
		{"", "    ", true},
		{"", "  \r\n  ", true},
		{"", "\014\012\011\013\015", true},
		{"", "\014\012\011\013 abc  \015", true},
		{"", "\f\n\t\v\r\f", true},
		{"", "x\n\t\t\t\t", true},
		{"", "\f\n\t  \n\n\n   \v\r\f", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasWhitespace(tt.str); got != tt.want {
				t.Errorf("HasWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
