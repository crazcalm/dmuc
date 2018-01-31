dmuc
====
[![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/dmuc)](https://goreportcard.com/report/github.com/crazcalm/dmuc)

Back Story:
-----------
dmuc is short for Discover My Unix Commands. This is a port of the original python version [python version](https://github.com/crazcalm/DiscoverMyUnixCommands)  to Golang

Install:
-------
* Clone repo
* go build/run dmuc.go

Usage:
------
	Usage of dmuc:
		-a	List files from both /usr/bin and /usr/local/bin directory
		-i string
			Filters putput based on if the content starts with the provided string
		-l	List files from /usr/local/bin directory
		-s string
			Filters output based on if the content includes the provided string
