// Package luhn is a set of tools for use with luhn algorithm
package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a given string passes the luhn algorithm checksum
func Valid(n string) bool {
	n = strings.Replace(n, " ", "", -1)
	runes := []rune(n)
	strlen := len(runes)
	if strlen <= 1 {
		return false
	}

	var factor, result int
	var char rune
	for i := len(runes) - 1; i >= 0; i-- {
		char = runes[i]
		if !unicode.IsNumber(char) {
			return false
		}
		if (strlen-i-1)%2 == 0 {
			factor = 1
		} else {
			factor = 2
		}

		num := int(char-'0') * factor
		if num <= 9 {
			result += num
		} else {
			result += num - 9
		}
	}
	return result%10 == 0
}
