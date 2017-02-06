package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var grep1 = flag.String("s", "", "Applies a grep 'char' filter to the output")
var grep2 = flag.String("i", "", "Apllies a grep 'string' filter to the output")

func grepCalls(paths string, output bytes.Buffer, grepString string, charGrep bool) string {
	args := strings.Split(paths, " ")
	var grepArg string
	var cmd *exec.Cmd

	if charGrep {
		grepArg = fmt.Sprintf("^%s", grepString)
	} else {
		grepArg = grepString
	}
	
	if len(args) == 1 {
		cmd = exec.Command("ls", args[0], "grep", grepArg)
	}else{
		cmd = exec.Command("ls", args[0], args[1], "grep", grepArg)
	}
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	return output.String()
}

func lsCall(paths string, output bytes.Buffer) string {
	args := strings.Split(paths, " ")
	var cmd *exec.Cmd
	if len(args) == 1 {
		cmd = exec.Command("ls", args[0])
	} else {
		cmd = exec.Command("ls", args[0], args[1])
	}
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	return output.String()
}

func printToScreen(content string) {
	fmt.Print(content)
}

func main() {
	cmd := exec.Command("ls", "/usr/bin/", "/usr/local/bin")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	printToScreen(out.String())

	flag.Parse()
	var out2 bytes.Buffer
	printToScreen(lsCall("/usr/bin", out2))
}
