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
	regexEmailPattern       = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	regexStrictEmailPattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
		`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
		`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
		`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	regexURLPattern = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@\-\/]))?`
	// RGB颜色
	regexRGBColorPattern = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	// 十六进制颜色
	regexHexColorPattern = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
)

var (
	RegexEmail       = regexp.MustCompile(regexEmailPattern)
	RegexStrictEmail = regexp.MustCompile(regexStrictEmailPattern)
	RegexURL         = regexp.MustCompile(regexURLPattern)
	RegexRGBColor    = regexp.MustCompile(regexRGBColorPattern)
	RegexHexColor    = regexp.MustCompile(regexHexColorPattern)
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
