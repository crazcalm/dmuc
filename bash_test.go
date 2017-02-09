package dmuc

import (
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
