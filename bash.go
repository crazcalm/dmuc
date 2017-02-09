package dmuc

import (
	"fmt"
)

// CreateBashCommand compiles the arguements needed for running the ls and grep commands
func CreateBashCommand(local, all *bool, grepStartsWith, grepIncludes *string) ([]string, []string) {
	lsArgs := []string{LS}
	if *all {
		lsArgs = append(lsArgs, USERBIN)
		lsArgs = append(lsArgs, USERLOCALBIN)
	} else if *local {
		lsArgs = append(lsArgs, USERLOCALBIN)
	} else {
		lsArgs = append(lsArgs, USERBIN)
	}

	grepArgs := []string{GREP}

	if *grepStartsWith != "" || *grepIncludes != "" {
		if *grepIncludes != "" {
			grepArgs = append(grepArgs, fmt.Sprintf("%s", *grepIncludes))
		} else {
			grepArgs = append(grepArgs, fmt.Sprintf("^%s", *grepStartsWith))
		}
	}

	return lsArgs, grepArgs
}
