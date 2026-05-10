package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toCamelCase(s string) string {
	words := splitWords(s)
	if len(words) == 0 {
		return s
	}
	result := strings.ToLower(words[0])
	for _, w := range words[1:] {
		if len(w) == 0 {
			continue
		}
		result += strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
	}
	return result
}

func toSnakeCase(s string) string {
	words := splitWords(s)
	lower := make([]string, 0, len(words))
	for _, w := range words {
		lower = append(lower, strings.ToLower(w))
	}
	return strings.Join(lower, "_")
}

func toCapitalCase(s string) string {
	words := splitWords(s)
	titled := make([]string, 0, len(words))
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		titled = append(titled, strings.ToUpper(w[:1])+strings.ToLower(w[1:]))
	}
	return strings.Join(titled, " ")
}

// splitWords splits a string on whitespace, underscores, hyphens and
// camelCase boundaries so that every transformation starts from clean tokens.
func splitWords(s string) []string {
	// First, insert a space before every upper-case letter that follows a
	// lower-case letter (camelCase → camel Case).
	var expanded strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) && unicode.IsLower(runes[i-1]) {
			expanded.WriteRune(' ')
		}
		expanded.WriteRune(r)
	}

	// Then split on any non-letter/non-digit character.
	return strings.FieldsFunc(expanded.String(), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func main() {
	upper := flag.Bool("upper", false, "convert args to UPPER CASE")
	camel := flag.Bool("camel", false, "convert args to camelCase")
	snake := flag.Bool("snake", false, "convert args to snake_case")
	capital := flag.Bool("capital", false, "convert args to Capital Case")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: test [--upper | --camel | --snake | --capital] <word> [word ...]")
		flag.Usage()
		os.Exit(1)
	}

	for _, arg := range args {
		switch {
		case *upper:
			fmt.Println(toUpperCase(arg))
		case *camel:
			fmt.Println(toCamelCase(arg))
		case *snake:
			fmt.Println(toSnakeCase(arg))
		case *capital:
			fmt.Println(toCapitalCase(arg))
		default:
			fmt.Println(arg)
		}
	}
}
