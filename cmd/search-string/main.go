package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/martinvw/go-search-string/cmd/search-string/internal"
	"log"
	"os"
	"sort"
	"strings"
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

	if file == "" || prefix == "" {
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
	beginIndex := strings.Index(haystack, prefix) + len(prefix)
	if beginIndex == -1 {
		return ""
	}
	endIndex := strings.Index(haystack[beginIndex:], postfix)
	return internal.Reverse(haystack[beginIndex : beginIndex+endIndex])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
