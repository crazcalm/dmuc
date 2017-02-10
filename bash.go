package dmuc

import (
	"bytes"
	"fmt"
	"os/exec"
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

// RunCommand runs the passed arugments in bash via exec.Command
func RunCommand(lsArgs []string, grepArgs []string, output bytes.Buffer) string {
	var lsCmd *exec.Cmd
	var grepCmd *exec.Cmd

	numOfArgs := len(lsArgs)

	if numOfArgs == 2 {
		lsCmd = exec.Command(lsArgs[0], lsArgs[1])
	} else if numOfArgs == 3 {
		lsCmd = exec.Command(lsArgs[0], lsArgs[1], lsArgs[2])
	}

	if len(grepArgs) == 2 {
		grepCmd = exec.Command(grepArgs[0], grepArgs[1])
		var err error

		grepCmd.Stdout = &output

		grepCmd.Stdin, err = lsCmd.StdoutPipe()
		grepCmd.Start()
		ErrorHandler(err)

		err = lsCmd.Start()
		ErrorHandler(err)

		err = lsCmd.Wait()
		ErrorHandler(err)

		// There is an error when the output is empty
		// So... ignore it?
		_ = grepCmd.Wait()
	} else {
		lsCmd.Stdout = &output
		err := lsCmd.Run()
		ErrorHandler(err)
	}
	return output.String()
}
