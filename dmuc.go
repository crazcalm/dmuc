package main

import (
	"flag"
	"github.com/crazcalm/dmuc/src"
)

var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
var i = flag.String("i", "", "Filters output based on if the content includes the provided string")
var s = flag.String("s", "", "Filters putput based on if the content starts with the provided string")

func main() {
	flag.Parse()
	dirs, filter, startFilter := dmuc.CreateCommand(l, a, s, i)
	dmuc.Run(dirs, filter, startFilter)
}
