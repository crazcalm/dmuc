# dmuc

[![Build Status](https://api.travis-ci.org/crazcalm/dmuc.svg?branch=master)](https://travis-ci.org/crazcalm/dmuc)     [![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/dmuc)](https://goreportcard.com/report/github.com/crazcalm/dmuc)     [![Coverage Status](https://coveralls.io/repos/github/crazcalm/dmuc/badge.svg?branch=master)](https://coveralls.io/github/crazcalm/dmuc?branch=master)

## Purpose:
This application allows you to list out the applications in the /usr/bin and /usr/local/bin directories. You may also use the "starts with" or "includes" filters to filter your results.


## Back Story:
dmuc is short for Discover My Unix Commands. This is a port of the original python version [python version](https://github.com/crazcalm/DiscoverMyUnixCommands)  to Golang

## Install:
* go get github/crazcalm/dmuc
* go build/run dmuc.go

## Usage:
	Usage of dmuc:
  	-a    List files from both /usr/bin and /usr/local/bin directory
  	-i string
          Filters output based on if the content includes the provided string
  	-l    List files from /usr/local/bin directory
  	-s string
          Filters output based on if the content starts with the provided string
