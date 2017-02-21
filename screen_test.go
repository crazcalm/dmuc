package dmuc

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintToScreen(t *testing.T) {
	var tests = []struct {
		input string
		dir   string
		want  []string
	}{
		{"", "", []string{"None"}},
		{"hello", "", []string{"hello"}},
		{"", "dir", []string{"None", "dir"}},
		{"hello", "dir", []string{"hello", "dir"}},
	}
	b := new(bytes.Buffer)
	var got string
	for _, test := range tests {
		PrintToScreen(b, test.input, test.dir)

		got = b.String()

		for _, item := range test.want {
			if strings.Contains(item, got) {
				t.Errorf("PrintToScreen(%s):  %s not in %s", test.input, test.want, got)
			}
		}
		b.Reset()
	}
}
