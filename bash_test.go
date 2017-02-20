package dmuc

import (
	"bytes"
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

func TestFilter(t *testing.T) {
	data := []string{"marcus", "allen", "willock"}

	var tests = []struct {
		content 	[]string
		filter 		string
		startsWith	bool
		answer		[]string
	}{
		{data, "", false, data},
		{data, "", true, data},
		{data, "a", false, []string{"marcus", "allen"}},
		{data, "a", true, []string{"allen"}},		
	}

	for _, test := range tests {
		output := Filter(test.content, test.filter, test.startsWith)

		if compare(output, test.answer) != true {
			t.Errorf("Filter(%v, %s, %t)\n output %v != %v", test.content, test.filter, test.startsWith, output, test.answer)
		}
	}
}

func TestLsCommand(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Failed to create a path variable for testing LSCommnad")
	}

	var tests = []struct {
		path	string
		files	[]string
	}{
		{path, []string{"bash.go", "bash_test.go", "screen.go", "screen_test.go", "constants.go"}},
	}

	for _, test := range tests {
		output := LsCommand(test.path)
		outputString := strings.Join(output, ", ")

		for _, file := range test.files {
			if strings.Contains(outputString, file) != true {
				t.Errorf("LsCommand(%s)\n output %s does not contain %s", test.path, outputString, file)
			}
		}
	}
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

		for _, value := range test.contains {
			if !strings.Contains(result, value) {
				t.Errorf("RunCommand(%v, %v, %v)\n%s in in %s", test.lsArgs, test.grepArgs, test.output, value, result)
			}
		}
	}
}

func TestCreateCommand(t *testing.T) {
	var tests = []struct {
		local		bool
		all			bool
		startsWith	string
		includes	string
		dirs		[]string
		filter		string
		startFilter	bool
	} {
		{false, false, "", "", []string{USERBIN}, "", false},
		{true, false, "", "", []string{USERLOCALBIN}, "", false},
		{true, true, "", "", []string{USERBIN, USERLOCALBIN}, "", false},
		{false, true, "", "", []string{USERBIN, USERLOCALBIN}, "", false},
		{false, false, "s", "", []string{USERBIN}, "s", true},
		{false, false, "", "i", []string{USERBIN}, "i", false},
		{false, false, "s", "i", []string{USERBIN}, "i", false},
	}

	for _, test :=range tests {
		dirs, filter, startFilter := CreateCommand(&test.local, &test.all, &test.startsWith, &test.includes)

		if compare(dirs, test.dirs) == false {
			t.Errorf("CreateCommand(%t, %t, %s, %s)\n dirs %v != %v", test.local, test.all, test.startsWith, test.includes, dirs, test.dirs)
		}

		if filter != test.filter {
			t.Errorf("CreateCommand(%t, %t, %s, %s)\n filter %s != %s", test.local, test.all, test.startsWith, test.includes, filter, test.filter)
		}

		if startFilter != test.startFilter {
			t.Errorf("CreateCommand(%t, %t, %s, %s)\n startsWith %t != %t", test.local, test.all, test.startsWith, test.includes, startFilter, test.startFilter)
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
