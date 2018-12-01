// Package grains gives you counts of grains on a chess board
package grains

import "errors"

type grainCounts struct {
	total  uint64
	square uint64
}

// Total returns the total count of grains ont he 64 cells of a chess board
func Total() uint64 {
	grainCounts, _ := countGrains(64)
	return grainCounts.total
}

// Square returns the count of grains on a single square of the chess board
func Square(n int) (uint64, error) {
	grainCounts, err := countGrains(n)
	return grainCounts.square, err
}

func countGrains(n int) (grainCounts, error) {
	if n <= 0 || n > 64 {
		return grainCounts{}, errors.New("bad input")
	}
	var total, square uint64 = 1, 1
	for i := 1; i < n; i++ {
		square = square * 2
		total += square
	}

	return grainCounts{total, square}, nil
}
