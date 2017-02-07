package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"log"
)

const (
	userBin = "/usr/bin/"
	userLocalBin = "/usr/local/bin/"
	pipe = "|"
	grep = "grep"
	LS = "ls"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var s = flag.String("s", "", "Applies a grep 'char' filter to the output")
var i = flag.String("i", "", "Apllies a grep 'string' filter to the output")

func createBashCommand(local *bool, all *bool, grepStartsWith *string, grepIncludes* string) []string {
	args := []string{LS}
	if *all{
		args = append(args, userBin)
		args = append(args, userLocalBin)
	}else if *local {
		args = append(args, userLocalBin)
	} else {
		args = append(args, userBin)
	}

	if *grepStartsWith != "" || *grepIncludes != ""{
		args = append(args, pipe)
		args = append(args, grep)
		if *grepStartsWith != ""{
			args = append(args, fmt.Sprintf("'^%s'", grepStartsWith))
		} else {
			args = append(args, fmt.Sprintf("'%s'", grepIncludes))
		}
	}
	return args
	
}

func errorHandler(err error){
	if err !=nil {
		log.Fatal(err)
	}
}

func runCommand(args []string, output bytes.Buffer) {
	var cmd *exec.Cmd
	numOfArgs := len(args)
	fmt.Printf("num of args: %d\n", numOfArgs)

	if numOfArgs == 2 {
		cmd = exec.Command(args[0], args[1])
	} else if numOfArgs == 3 {
		cmd = exec.Command(args[0], args[1], args[2])
	} else if numOfArgs == 4 {
		cmd = exec.Command(args[0], args[1], args[2], args[3])
	} else if numOfArgs == 5 {
		cmd = exec.Command(args[0], args[1], args[2], args[3], args[4])
	} else if numOfArgs == 6 {
		cmd = exec.Command(args[0], args[1], args[2], args[3], args[4], args[5])
	}
	fmt.Println(args)
	cmd.Stdout = &output

	err := cmd.Run()

	fmt.Print("output: ")
	fmt.Println(output.String())

	errorHandler(err)
}

func printToScreen(content string) {
	fmt.Println(content)
	fmt.Println("over")
}

func main() {
	flag.Parse()

	var output bytes.Buffer
	args := createBashCommand(l,a,s,i)
	runCommand(args, output)
	printToScreen(output.String())
}
