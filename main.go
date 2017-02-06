package main

import (
	"flag"
	"fmt"
	"strings"
	"os/exec"
	"os"
	"bytes"
)


var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")

func lsCall(paths string, output bytes.Buffer){
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
	fmt.Print(output.String())
}


func main(){
	cmd := exec.Command("ls", "/usr/bin/", "/usr/local/bin")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	fmt.Print(out.String())
	
	flag.Parse()
	var out2  bytes.Buffer
	lsCall("/usr/bin", out2)
}
