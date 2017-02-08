package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

const (
	userBin      = "/usr/bin/"
	userLocalBin = "/usr/local/bin/"
	grep         = "grep"
	LS           = "ls"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var s = flag.String("s", "", "Applies a grep 'char' filter to the output")
var i = flag.String("i", "", "Apllies a grep 'string' filter to the output")

func createBashCommand(local *bool, all *bool, grepStartsWith *string, grepIncludes *string) ([]string, []string) {
	lsArgs := []string{LS}
	if *all {
		lsArgs = append(lsArgs, userBin)
		lsArgs = append(lsArgs, userLocalBin)
	} else if *local {
		lsArgs = append(lsArgs, userLocalBin)
	} else {
		lsArgs = append(lsArgs, userBin)
	}

	grepArgs := []string{grep}

	if *grepStartsWith != "" || *grepIncludes != "" {
		if *grepStartsWith != "" {
			grepArgs = append(grepArgs, fmt.Sprintf("^%s", *grepStartsWith))
		} else {
			grepArgs = append(grepArgs, fmt.Sprintf("%s", *grepIncludes))
		}
	}
	return lsArgs, grepArgs

}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
		errorHandler(err)

		err = lsCmd.Start()
		errorHandler(err)

		err = lsCmd.Wait()
		errorHandler(err)

		// There is an error when the output is empty
		//So... ignore it?
		_ = grepCmd.Wait()
	} else {
		lsCmd.Stdout = &output
		err := lsCmd.Run()
		errorHandler(err)
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
