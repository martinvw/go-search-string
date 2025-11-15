package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/martinvw/go-search-string/cmd/search-string/internal"
)

var file string
var needlePrefix string
var prefix string
var postfix string

func init() {
	flag.StringVar(&file, "f", "", "filename containing filter values")
	flag.StringVar(&prefix, "locator-prefix", "", "prefix used to locate/extract the target string")
	flag.StringVar(&postfix, "locator-postfix", "", "postfix used to locate/extract the target string")
	flag.StringVar(&needlePrefix, "needle-prefix", "", "prefix to apply to every search string")
}

func main() {
	flag.Parse()

	if file == "" {
		flag.Usage()
		os.Exit(1)
	}

	searchStrings := internal.LoadFilterFile(file, needlePrefix)

	processStdIn(searchStrings)
}

func processStdIn(searchStrings []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		haystack := scanner.Text()
		potentialNeedle := findPotentialNeedle(haystack, prefix, postfix)
		needleLength := len(potentialNeedle)

		i := sort.Search(len(searchStrings), func(i int) bool {
			minLength := min(len(searchStrings[i]), needleLength)
			return searchStrings[i][:minLength] >= potentialNeedle[:minLength]
		})

		if i < len(searchStrings) && strings.HasPrefix(potentialNeedle, searchStrings[i]) {
			fmt.Println(haystack)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func findPotentialNeedle(haystack string, prefix string, postfix string) string {
	beginIndex := 0
	if prefix == "" {
		beginIndex = 0
	} else {
		beginIndex = strings.Index(haystack, prefix) + len(prefix)
		if beginIndex == -1 {
			return ""
		}
	}
	endIndex := 0
	if postfix == "" {
		endIndex = len(haystack) - beginIndex
	} else {
		endIndex = strings.Index(haystack[beginIndex:], postfix)
		// if we cannot find postfix, we have a mismatch
		if endIndex == -1 {
			return ""
		}
	}
	return internal.Reverse(haystack[beginIndex : beginIndex+endIndex])
}
