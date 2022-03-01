package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// LoadFilterFile Load all filters from the given filter file
// all filters will be pre/postfixed with the given prefix / postfix
func LoadFilterFile(filename string, prefix string) []string {
	searchStrings := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(file)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		searchStrings = append(searchStrings, Reverse(prefix+fileScanner.Text()))
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(searchStrings)

	return searchStrings
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
