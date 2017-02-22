dmuc
====

Back Story:
-----------
	dmuc is short for Discover My Unix Commands. This is a port of
	the original python version to Golang

Install:
-------
* go get github.com/crazcalm/dmuc
* Copy the below fale (dmuc.go):
	package main

	import (
	        "flag"
	        "github.com/crazcalm/dmuc"
	)

	var l = flag.Bool("l", false, "List files from /usr/local/bin directory")
	var a = flag.Bool("a", false, "List files from both /usr/bin and /usr/local/bin directory")
	var s = flag.String("s", "", "Filters output based on if the content includes the provided string")
	var i = flag.String("i", "", "Filters putput based on if the content starts with the provided string")

	func main() {
	        flag.Parse()
	        dirs, filter, startFilter := dmuc.CreateCommand(l, a, s, i)
	        dmuc.Run(dirs, filter, startFilter)
	}
* go build dmuc.go

Usage:
------
	Usage of ./main2:
		-a	List files from both /usr/bin and /usr/local/bin directory
		-i string
			Filters putput based on if the content starts with the provided string
		-l	List files from /usr/local/bin directory
		-s string
			Filters output based on if the content includes the provided string
