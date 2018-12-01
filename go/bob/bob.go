// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isNoLetters(remark) {
		return "Whatever."
	}
	if isForcefulQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isLoud(remark) {
		return "Whoa, chill out!"
	}
	if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}

func isNoLetters(remark string) bool {
	// return regexp.MustCompile(`^[0-9]+$`).MatchString(remark)
	return regexp.MustCompile(`(\d+)(,\s*\d+)*`).MatchString(remark)
}

func isForcefulQuestion(remark string) bool {
	return isLoud(remark) && isQuestion(remark)
}

func isExcited(remark string) bool {
	return strings.HasSuffix(remark, "!")
}

func isLoud(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
