package dmuc

import (
	"fmt"
	"io"
	"log"
)

// PrintToScreen prints the results to screen
func PrintToScreen(w io.Writer, content, dir string) {
	_, err := fmt.Fprintf(w, "%s:\n", dir)
	if err != nil {
		log.Fatal(fmt.Errorf("PrintToScreen error: %s", err.Error()))
	}
	if content == "" {
		_, err = fmt.Fprintln(w, "None")
		if err != nil {
			log.Fatal(fmt.Errorf("PrintToScreen error: %s", err.Error()))
		}
	} else {
		_, err = fmt.Fprintln(w, content)
		if err != nil {
			log.Fatal(fmt.Errorf("PrintToScreen error: %s", err.Error()))
		}
	}
}
