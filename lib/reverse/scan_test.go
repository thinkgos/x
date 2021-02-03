package reverse

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	var tests = []struct {
		text   string
		split  bufio.SplitFunc
		tokens []string
	}{
		{
			text: `
hello one two three four five six seven
there

you one two three four five six seven` + "\r" + `
x

`,
			split: bufio.ScanLines,
			tokens: []string{
				"",
				"x",
				"you one two three four five six seven",
				"",
				"there",
				"hello one two three four five six seven",
				"",
			},
		},
		{
			text: `
hello one two three four five six seven
there

you one two three four five six seven` + "\r" + `
x
`,
			split: bufio.ScanWords,
			tokens: []string{
				"x",
				"seven",
				"six",
				"five",
				"four",
				"three",
				"two",
				"one",
				"you",
				"there",
				"seven",
				"six",
				"five",
				"four",
				"three",
				"two",
				"one",
				"hello",
			},
		},
		// {// TODO: bug 空格并没有略过
		// 	text:   "one" + strings.Repeat(" ", 50),
		// 	split:  bufio.ScanWords,
		// 	tokens: []string{"one"},
		// },
		// {
		// 	text:   "1234567890     one        ",
		// 	split:  bufio.ScanWords,
		// 	tokens: []string{"one", "1234567890"},
		// },
		{
			text:   "one",
			split:  bufio.ScanWords,
			tokens: []string{"one"},
		},
		{
			text:  "",
			split: bufio.ScanWords,
		},
	}

	for i, test := range tests {
		b := NewScanner(strings.NewReader(test.text))
		b.buf = make([]byte, 10)
		b.Split(test.split)
		var got []string
		for b.Scan() {
			got = append(got, b.Text())
		}
		if b.Err() != nil {
			t.Fatalf("error after scan: %v", b.Err())
		}
		if !reflect.DeepEqual(got, test.tokens) {
			t.Fatalf("test %d: token mismatch; got %q want %q", i, got, test.tokens)
		}
	}
}
