package dmuc

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"testing"
)

func compare(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRunCommand(t *testing.T) {
	//If I could get the pwd dir via go, that would be ideal.
	//That way I could compare the dir name to the ls output.
	//I should use "path/filepath"
	var output bytes.Buffer
	path, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("failed to create a path variable for testing the RunCommand")
	}

	var tests = []struct {
		lsArgs   []string
		grepArgs []string
		output   bytes.Buffer
		contains []string
	}{
		{[]string{LS, path}, []string{GREP}, output, []string{"bash_test.go", "bash.go"}},
		{[]string{LS, path, path}, []string{GREP}, output, []string{"/dmuc"}},
		{[]string{LS, path}, []string{GREP, "test"}, output, []string{"bash_test.go", "screen_test.go"}},
		{[]string{LS, path}, []string{GREP, "^t"}, output, []string{"constants.go", "bash_test.go"}},
	}

	for _, test := range tests {
		result := RunCommand(test.lsArgs, test.grepArgs, test.output)

		fmt.Println(result)

		for _, value := range test.contains {
			if !strings.Contains(result, value) {
				t.Errorf("RunCommand(%v, %v, %v)\n%s in in %s", test.lsArgs, test.grepArgs, test.output, value, result)
			}
		}
	}
}

func TestCreateBashCommand(t *testing.T) {
	var tests = []struct {
		local          bool
		all            bool
		grepStartsWith string
		grepIncludes   string
		lsArgs         []string
		grepArgs       []string
	}{
		{false, false, "", "", []string{LS, USERBIN}, []string{GREP}},
		{true, false, "", "", []string{LS, USERLOCALBIN}, []string{GREP}},
		{true, true, "", "", []string{LS, USERBIN, USERLOCALBIN}, []string{GREP}},
		{false, true, "", "", []string{LS, USERBIN, USERLOCALBIN}, []string{GREP}},
		{false, false, "s", "", []string{LS, USERBIN}, []string{GREP, "^s"}},
		{false, false, "", "i", []string{LS, USERBIN}, []string{GREP, "i"}},
		{false, false, "s", "i", []string{LS, USERBIN}, []string{GREP, "i"}},
	}

	for _, test := range tests {
		gotLs, gotGrep := CreateBashCommand(&test.local, &test.all, &test.grepStartsWith, &test.grepIncludes)

		if compare(gotLs, test.lsArgs) == false {
			t.Errorf("CreateBashCommand(%t, %t, %s, %s)\nlsArgs: %v != %v", test.local, test.all, test.grepStartsWith, test.grepIncludes, gotLs, test.lsArgs)
		}
		if compare(gotGrep, test.grepArgs) == false {
			t.Errorf("CreateBashCommand(%t, %t, %s, %s)\ngrepArgs: %v != %v", test.local, test.all, test.grepStartsWith, test.grepIncludes, gotGrep, test.grepArgs)
		}
	}
}
