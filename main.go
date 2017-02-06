package main

import (
	"flag"
	"fmt"
	"strings"
	"os/exec"
	"os"
	"bytes"
)


var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")

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
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n{
		fmt.Println()
	}
}
