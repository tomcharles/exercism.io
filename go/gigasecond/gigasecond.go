// Package gigasecond tells you some good stuff about times and gigaseconds
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds 1,000,000,000 seconds to the given time
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Second * time.Duration(math.Pow(10, 9))
	return t.Add(gigasecond)
}
