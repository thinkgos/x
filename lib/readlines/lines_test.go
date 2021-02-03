package readlines

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	var tests = []struct {
		input     string
		bufioSize int
		maxSize   int
		want      []string
	}{
		{
			input: `
one
two
three
`[1:],
			maxSize: 10,
			want:    []string{"one", "two", "three"},
		},
		{
			input: `
01234567890123456789
01234567890123456
0123
`[1:],
			bufioSize: 16,
			maxSize:   18,
			want:      []string{"012345678901234567", "01234567890123456", "0123"},
		},
		{
			input: `
0123456789abcdefghijklmnopqrstuvwxyz!@#$%%^&*
0123456789abcdefghijklmnopqrstuvwxyz!@#$%%^&*
`[1:],
			maxSize: 10,
			want:    []string{"0123456789", "0123456789"},
		},
		{
			input:   "oneline",
			maxSize: 20,
			want:    []string{"oneline"},
		},
		{
			input:   "\n\n",
			maxSize: 20,
			want:    []string{"", ""},
		},
		{
			input:   `Peppé`,
			maxSize: 5,
			want:    []string{"Pepp"},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var r io.Reader = strings.NewReader(test.input)
			if test.bufioSize > 0 {
				r = bufio.NewReaderSize(r, test.bufioSize)
			}
			var got []string
			err := Iter(r, test.maxSize, func(line []byte) error {
				got = append(got, string(line))
				return nil
			})
			assert.Nil(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func ExampleIter() {
	input := `one two
three
four

five`
	r := strings.NewReader(input)
	Iter(r, 1024*1024, func(line []byte) error {
		fmt.Printf("%q\n", line)
		return nil
	})
	// Output:
	// "one two"
	// "three"
	// "four"
	// ""
	// "five"
}
