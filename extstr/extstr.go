package extstr

import (
	"math/rand"
	"strconv"
	"strings"
)

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []int64, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(elems[0], 10)
	}
	strElems := make([]string, 0, len(elems))
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		v := strconv.FormatInt(elems[i], 10)
		strElems = Append(strElems, v)
		n += len(v)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strElems[0])
	for _, s := range strElems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

// JoinInt concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinInt(elems []int, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(elems[0])
	}
	strElems := make([]string, 0, len(elems))
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		v := strconv.Itoa(elems[i])
		strElems = Append(strElems, v)
		n += len(v)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strElems[0])
	for _, s := range strElems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

// Split Split slices s into all substrings separated by sep and returns a slice of
// the int64 between those separators.
func Split(s, sep string) []int64 {
	if s == "" {
		return []int64{}
	}

	ss := strings.Split(s, sep)
	res := make([]int64, 0, len(ss))
	for i := 0; i < len(ss); i++ {
		v, err := strconv.ParseInt(strings.TrimSpace(ss[i]), 10, 64)
		if err != nil {
			continue
		}
		res = append(res, v)
	}
	return res
}

// SplitInt Split slices s into all substrings separated by sep and returns a slice of
// the int between those separators.
func SplitInt(s, sep string) []int {
	if s == "" {
		return []int{}
	}

	ss := strings.Split(s, sep)
	res := make([]int, 0, len(ss))
	for i := 0; i < len(ss); i++ {
		v, err := strconv.Atoi(strings.TrimSpace(ss[i]))
		if err != nil {
			continue
		}
		res = append(res, v)
	}
	return res
}

// ShuffleString pseudo-randomizes the order of elements using the default Source.
func Shuffle(str string) string {
	runes := []rune(str)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
