package normalize

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EOL(t *testing.T) {
	data1 := []string{
		"",
		"This text starts with empty lines",
		"another",
		"",
		"",
		"",
		"Some other empty lines in the middle",
		"more.",
		"And more.",
		"Ends with empty lines too.",
		"",
		"",
		"",
	}

	data2 := []string{
		"This text does not start with empty lines",
		"another",
		"",
		"",
		"",
		"Some other empty lines in the middle",
		"more.",
		"And more.",
		"Ends without EOLtoo.",
	}

	buildEOLData := func(data []string, eol string) []byte {
		return []byte(strings.Join(data, eol))
	}

	dos := buildEOLData(data1, "\r\n")
	unix := buildEOLData(data1, "\n")
	mac := buildEOLData(data1, "\r")

	assert.Equal(t, unix, EOL(dos))
	assert.Equal(t, unix, EOL(mac))
	assert.Equal(t, unix, EOL(unix))

	dos = buildEOLData(data2, "\r\n")
	unix = buildEOLData(data2, "\n")
	mac = buildEOLData(data2, "\r")

	assert.Equal(t, unix, EOL(dos))
	assert.Equal(t, unix, EOL(mac))
	assert.Equal(t, unix, EOL(unix))

	assert.Equal(t, []byte("one liner"), EOL([]byte("one liner")))
	assert.Equal(t, []byte("\n"), EOL([]byte("\n")))
	assert.Equal(t, []byte("\ntwo liner"), EOL([]byte("\ntwo liner")))
	assert.Equal(t, []byte("two liner\n"), EOL([]byte("two liner\n")))
	assert.Equal(t, []byte{}, EOL([]byte{}))

	assert.Equal(t, []byte("mix\nand\nmatch\n."), EOL([]byte("mix\r\nand\rmatch\n.")))
}
