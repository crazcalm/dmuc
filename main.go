package dmuc

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var s = flag.String("s", "", "Applies a grep 'char' filter to the output")
var i = flag.String("i", "", "Apllies a grep 'string' filter to the output")

func createBashCommand(local *bool, all *bool, grepStartsWith *string, grepIncludes *string) ([]string, []string) {
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
		if *grepStartsWith != "" {
			grepArgs = append(grepArgs, fmt.Sprintf("^%s", *grepStartsWith))
		} else {
			grepArgs = append(grepArgs, fmt.Sprintf("%s", *grepIncludes))
		}
	}
	return lsArgs, grepArgs

}


func runCommand(lsArgs []string, grepArgs []string, output bytes.Buffer) string {
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
		//So... ignore it?
		_ = grepCmd.Wait()
	} else {
		lsCmd.Stdout = &output
		err := lsCmd.Run()
		ErrorHandler(err)
	}

	return output.String()
}

func printToScreen(content string) {
	fmt.Println("See Results Below: ")
	if content == "" {
		fmt.Println("None")
	} else {
		fmt.Println(content)
	}
}

func main() {
	flag.Parse()

	var output bytes.Buffer
	lsArgs, grepArgs := createBashCommand(l, a, s, i)
	result := runCommand(lsArgs, grepArgs, output)
	printToScreen(result)
}
