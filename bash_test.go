package dmuc

import (
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
	dataSorted := []string{"allen", "marcus", "willock"}

	var tests = []struct {
		content    []string
		filter     string
		startsWith bool
		answer     []string
	}{
		{data, "", false, dataSorted},
		{data, "", true, dataSorted},
		{data, "a", false, []string{"allen", "marcus"}},
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
		path  string
		files []string
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

func TestCreateCommand(t *testing.T) {
	var tests = []struct {
		local       bool
		all         bool
		startsWith  string
		includes    string
		dirs        []string
		filter      string
		startFilter bool
	}{
		{false, false, "", "", []string{USERBIN}, "", false},
		{true, false, "", "", []string{USERLOCALBIN}, "", false},
		{true, true, "", "", []string{USERBIN, USERLOCALBIN}, "", false},
		{false, true, "", "", []string{USERBIN, USERLOCALBIN}, "", false},
		{false, false, "s", "", []string{USERBIN}, "s", true},
		{false, false, "", "i", []string{USERBIN}, "i", false},
		{false, false, "s", "i", []string{USERBIN}, "i", false},
	}

	for _, test := range tests {
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
