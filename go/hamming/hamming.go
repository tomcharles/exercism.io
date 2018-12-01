// Package hamming calculates the hamming between two DNA strands
package hamming

import "errors"

// Distance gives a numeric value representing the hamming stance between two strands or an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("input length does not match")
	}
	distance := 0
	for idx, char := range a {
		if rune(char) != rune(b[idx]) {
			distance = distance + 1
		}
	}
	return distance, nil
}
