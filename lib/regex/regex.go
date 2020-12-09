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
	"regexp"
	"strings"
)

const (
	patternEmail       = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	patternStrictEmail = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
		`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
		`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
		`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	patternURL = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@\-\/]))?`
	// RGB颜色
	patternRGBColor = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	// 十六进制颜色
	patternRegexHexColor = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
	// 字母或数字
	patternLetterNumeric = `^[a-zA-Z0-9]+$`
	// 全中文
	patternChinese = "^[\u4e00-\u9fa5]+$"
	// 中文名
	patternChineseName = "^[\u4e00-\u9fa5][.•·\u4e00-\u9fa5]{0,30}[\u4e00-\u9fa5]$"
	// 词语,不以下划线开头的中文、英文、数字、下划线、空格
	patternWord = "^[a-zA-Z0-9\u4e00-\u9fa5][a-zA-Z0-9_\u4e00-\u9fa5]+$"
	// 全空白字符
	patternWhitespaceAll = "^[[:space:]]+$"

	// 含有空白字符
	patternWhitespaceHas = ".*[[:space:]]"
)

var (
	RegexEmail         = regexp.MustCompile(patternEmail)
	RegexStrictEmail   = regexp.MustCompile(patternStrictEmail)
	RegexURL           = regexp.MustCompile(patternURL)
	RegexRGBColor      = regexp.MustCompile(patternRGBColor)
	RegexHexColor      = regexp.MustCompile(patternRegexHexColor)
	RegexAlphaNumeric  = regexp.MustCompile(patternLetterNumeric)
	RegexChinese       = regexp.MustCompile(patternChinese)
	RegexChineseName   = regexp.MustCompile(patternChineseName)
	RegexWord          = regexp.MustCompile(patternWord)
	RegexWhitespaceAll = regexp.MustCompile(patternWhitespaceAll)
	RegexWhitespaceHas = regexp.MustCompile(patternWhitespaceHas)
)

// IsEmail validates string is an email address, if not return false
// basically validation can match 99% cases
func IsEmail(s string) bool {
	return s != "" && RegexEmail.MatchString(s)
}

// IsEmailRFC validates string is an email address, if not return false
// this validation omits RFC 2822
func IsEmailRFC(email string) bool {
	return email != "" && RegexStrictEmail.MatchString(email)
}

// IsURL validates string is a url link, if not return false
// simple validation can match 99% cases
func IsURL(s string) bool {
	return s != "" && RegexURL.MatchString(s)
}

// IsRGBColor 是否是rgb颜色格式
// rgb(0,31, 255)
func IsRGBColor(s string) bool {
	return s != "" && RegexRGBColor.MatchString(s)
}

// IsHexColor 检查是否十六进制颜色,并返回带"#"的修正值.
func IsHexColor(s string) (string, bool) {
	ok := s != "" && RegexHexColor.MatchString(s)
	if ok {
		s = strings.ToUpper(s)
		if !strings.ContainsRune(s, '#') {
			s = "#" + s
		}
	}
	return s, ok
}

// IsAlphaNumeric 是否字母或数字.
func IsAlphaNumeric(s string) bool {
	return s != "" && RegexAlphaNumeric.MatchString(s)
}

// IsWord 是否词语(不以下划线开头的中文、英文、数字、下划线、空格).
func IsWord(str string) bool {
	return str != "" && RegexWord.MatchString(str)
}

// IsChinese 字符串是否全部中文.
func IsChinese(str string) bool {
	return str != "" && RegexChinese.MatchString(str)
}

// IsChineseName 字符串是否中文名.
func IsChineseName(str string) bool {
	return str != "" && RegexChineseName.MatchString(str)
}

// IsWhitespaces 是否全部空白字符,不包括空字符串.
func IsWhitespaces(str string) bool {
	return str != "" && RegexWhitespaceAll.MatchString(str)
}

func HasWhitespace(str string) bool {
	return str != "" && RegexWhitespaceHas.MatchString(str)
}
