# Go searching function

[![Release](https://img.shields.io/github/release/martinvw/go-search-string.svg?style=flat-square)](https://github.com/martinvw/go-search-string/releases/latest) 
[![PkgGoDev](https://pkg.go.dev/badge/github.com/martinvw/go-search-string)](https://pkg.go.dev/github.com/martinvw/go-search-string) 
[![Go Report Card](https://goreportcard.com/badge/github.com/martinvw/go-search-string?style=flat-square)](https://goreportcard.com/report/github.com/martinvw/go-search-string) 

Search for a list of strings in std-in and output all matches.

## Installation

```bash
go get -u github.com/martinvw/go-search-string/cmd/search-string
```

## Usage

```bash
Usage of search_string:
  -f string
        filename containing filter values
  -locator-postfix string
        postfix used to locate/extract the target string
  -locator-prefix string
        prefix used to locate/extract the target string
  -needle-prefix string
        prefix to apply to every search string
```

## Release & Downloads

Releases and binary builds can be found at [GitHub / Releases](https://github.com/martinvw/go-search-string/releases/)

## Examples

To find a set of subdomains in a big file, the usage could be as follows:

```bash
cat "big-haystack.txt" | search-string -f "needles.txt" --locator-prefix 'name":"' --locator-postfix "\"" --needle-prefix "."

# Or using compressed input & pigz
pigz -dc "big-haystack.gz" | search-string -f "needles.txt" --locator-prefix 'name":"' --locator-postfix "\"" --needle-prefix "."

# Or combining the above with GNU parallel
pigz -dc "big-haystack.gz" | parallel --pipe --block 100M -q search-string -f "$input" --locator-prefix 'name":"' --locator-postfix "\"" --needle-prefix "."
```