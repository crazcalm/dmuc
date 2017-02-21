package dmuc

import (
	"fmt"
	"io"
)

// PrintToScreen prints the results to screen
func PrintToScreen(w io.Writer, content, dir string) {
	fmt.Fprintf(w, "%s:\n", dir)
	if content == "" {
		fmt.Fprintln(w, "None")
	} else {
		fmt.Fprintln(w, content)
	}
}
