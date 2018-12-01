// Package raindrops gives you vital information about numbers => raindrops
package raindrops

import "strconv"

var factorSounds = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert takes a number of raindrops and makes sounds out of them
func Convert(num int) (out string) {
	out = ""
	for _, factorSound := range factorSounds {
		if num%factorSound.factor == 0 {
			out += factorSound.sound
		}
	}
	if out == "" {
		out = strconv.Itoa(num)
	}
	return
}
