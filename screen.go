package dmuc

import (
	"fmt"
	"io"
)

// PrintToScreen prints the results to screen
func PrintToScreen(w io.Writer, content string) {
	fmt.Fprintln(w, "See Results Below: ")
	if content == "" {
		fmt.Fprintln(w, "None")
	} else {
		fmt.Fprintln(w, content)
	}
}