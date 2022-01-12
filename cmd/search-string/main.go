package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/martinvw/go-search-string/cmd/search-string/internal"
	"log"
	"os"
	"strings"
)

var file string
var prefix string
var postfix string

func init() {
	flag.StringVar(&file, "f", "", "filename containing filter values")
	flag.StringVar(&prefix, "prefix", "", "postfix to apply to every search string")
	flag.StringVar(&postfix, "postfix", "", "postfix to apply to every search string")
}

func main() {
	flag.Parse()

	if file == "" {
		flag.Usage()
		os.Exit(1)
	}

	searchStrings := internal.LoadFilterFile(file, prefix, postfix)

	processStdIn(searchStrings)
}

func processStdIn(searchStrings []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		haystack := scanner.Text()
		for _, value := range searchStrings {
			if strings.Contains(haystack, value) {
				fmt.Println(haystack)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
