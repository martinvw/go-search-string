# Go searching function

Search for a list of strings in std-in and output all matches.

## Usage

```bash
Usage of search_string:
  -f string
        filename containing filter values
  --postfix string
        postfix to apply to every search string
  --prefix string
        postfix to apply to every search string
```

## Release & Downloads

Releases and binary builds can be found at [GitHub / Releases](https://github.com/martinvw/go-search-string/releases/)

## Examples

To find a set of subdomains in a big file, the usage could be as follows:

```bash
cat "big-haystack.txt" | search-string -f "needles.txt" --prefix "." --postfix "\""

# Or using compressed input & pigz
pigz -dc "big-haystack.gz" | search-string -f "needles.txt" --prefix "." --postfix "\""

# Or combining the above with GNU parallel
pigz -dc "big-haystack.gz" | parallel --pipe --block 100M -q search-string -f "$input" --prefix "." --postfix "\""
```