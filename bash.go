package dmuc

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"os"
	"log"
)

func LsCommand(dirName string) []string {
        //learn how to check if it is a directory
        stat, err := os.Stat(dirName)
        if err != nil {
                log.Fatal(err)
        }
        if stat.IsDir() != true {
                log.Fatalf("%s is not a directory", dirName)
        }

        f, err := os.Open(dirName)
        if err != nil {
                log.Fatal(err)
        }

        names, err := f.Readdirnames(-1)
        if err != nil {
                log.Fatal(err)
        }

        return names
}

func Filter(content []string, filter string, startsWith bool) []string{
        var result []string

        // Starts with filter
        if startsWith == true {
                for _, item := range content {
                        if strings.HasPrefix(item, filter) {
                                result = append(result, item)
                        }
                }
        } else { // Includes filter
                for _, item := range content {
                        if strings.Contains(item, filter) {
                                result = append(result, item)
                        }
                }
        }
        return result
}

func CreateCommand(local, all *bool, startsWith, includes *string) ([]string, string, bool) {
	var dirs  []string
	if *all {
		dirs = append(dirs, USERBIN)
		dirs = append(dirs, USERLOCALBIN)
	} else if *local {
		dirs = append(dirs, USERLOCALBIN)
	} else {
		dirs = append(dirs, USERBIN)
	}

	var filter string
	var startFilter bool

	if *startsWith != "" || *includes != "" {
		if *includes != "" {
			filter = *includes
		} else {
			filter = *startsWith
			startFilter = true
		}
	}
	return dirs, filter, startFilter
}

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
