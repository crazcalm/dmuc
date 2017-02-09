package dmuc

import(
	"testing"
	"bytes"
)

func TestPrintToScreen(t *testing.T){
	var tests = []struct{
		input string
		want string
	}{
		{"", "None"},
		{"hello", "hello"},
	}
	b := new(bytes.Buffer)
	var got string
	for _, test := range tests {
		PrintToScreen(b, test.input)

		got = b.String()
		
		if got != test.want{
			t.Errorf("PrintToScreen(%s) = %s", test.input, got)
		}
	}
}
