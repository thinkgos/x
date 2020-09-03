package extstr

import (
	"strings"
)

// Append appends string to slice with no duplicates.
// 追加字符串,无重复项
func Append(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// Delete 删除string切片中的,第一个出现的指定元素
func Delete(s []string, e string) []string {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteAll 删除string切片中的 所有出现的指示元素
func DeleteAll(s []string, e string) []string {
	if s == nil {
		return nil
	}

	tmpS := make([]string, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// Compare compares two 'string' type slices.
// It returns true if elements and order are both the same.
// 比较两个字符串切片,要求元素和顺序都一致才返回true
func Compare(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareU compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
// 比较两个字符串切片,要求元素一致,且忽略顺序,一致返回true
func CompareU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}

	return len(s2) == 0
}

// Contains returns true if the string exists in given slice
// 字符串切片是否含有指定的元素,大小写敏感
func Contains(sl []string, str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}

// ContainsFold returns true if the string exists in given slice, ignore case.
// 字符串切片是否含有指定的元素,忽略大小写
func ContainsFold(sl []string, str string) bool {
	for _, s := range sl {
		if strings.EqualFold(s, str) {
			return true
		}
	}
	return false
}

// Unique takes an input slice of strings and
// returns a new slice of strings without duplicate values.
func Unique(a []string) []string {
	l := len(a)
	if l <= 1 {
		return a
	}

	m := make(map[string]struct{}, l)
	r := make([]string, 0, l)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// Reverse a utf8 encoded string.
func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
