package dmuc

import (
	"log"
)

// ErrorHandler is a function that checks error values, logs errors
// and, if needed, exits the program
func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
