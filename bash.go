package dmuc

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//LsCommand returns the items in a directory
func LsCommand(dirName string) []string {
	//learn how to check if it is a directory
	stat, err := os.Stat(dirName)
	if err != nil {
		log.Fatal(err)
	}
	if stat.IsDir() != true {
		log.Fatalf("%s is not a directory", dirName)
	}

	f, err := os.Open(dirName)
	if err != nil {
		log.Fatal(err)
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	return names
}

//Filter uses a string to filter a slice of strings.
//You can filter by 'include' and 'startsWith'
func Filter(content []string, filter string, startsWith bool) []string {
	var result []string

	// Starts with filter
	if startsWith == true {
		for _, item := range content {
			if strings.HasPrefix(item, filter) {
				result = append(result, item)
			}
		}
	} else { // Includes filter
		for _, item := range content {
			if strings.Contains(item, filter) {
				result = append(result, item)
			}
		}
	}
	sort.Strings(result)
	return result
}

//CreateCommand formats the flag arguements into a state that is usable by other
//functions in this package
func CreateCommand(local, all *bool, startsWith, includes *string) ([]string, string, bool) {
	var dirs []string
	if *all {
		dirs = append(dirs, USERBIN)
		dirs = append(dirs, USERLOCALBIN)
	} else if *local {
		dirs = append(dirs, USERLOCALBIN)
	} else {
		dirs = append(dirs, USERBIN)
	}

	var filter string
	var startFilter bool

	if *startsWith != "" || *includes != "" {
		if *includes != "" {
			filter = *includes
		} else {
			filter = *startsWith
			startFilter = true
		}
	}
	return dirs, filter, startFilter
}

//Run runs the program
func Run(dirs []string, filter string, startFilter bool) {
	for _, dir := range dirs {
		items := LsCommand(dir)
		results := Filter(items, filter, startFilter)
		PrintToScreen(os.Stdout, strings.Join(results, "\n"), dir)
		fmt.Printf("\n")
	}
}
