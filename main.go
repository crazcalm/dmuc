package dmuc

import (
	"bytes"
	"flag"
	"os"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var s = flag.String("s", "", "Applies a grep 'char' filter to the output")
var i = flag.String("i", "", "Apllies a grep 'string' filter to the output")

func main() {
	flag.Parse()

	var output bytes.Buffer
	lsArgs, grepArgs := CreateBashCommand(l, a, s, i)
	result := RunCommand(lsArgs, grepArgs, output)
	PrintToScreen(os.Stdout, result)
}
