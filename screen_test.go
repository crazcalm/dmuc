package dmuc

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintToScreen(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", "None"},
		{"hello", "hello"},
	}
	b := new(bytes.Buffer)
	var got string
	for _, test := range tests {
		PrintToScreen(b, test.input)

		got = b.String()

		if strings.Contains(test.want, got) {
			t.Errorf("PrintToScreen(%s):  %s not in %s", test.input, test.want, got)
		}
		b.Reset()
	}
}
