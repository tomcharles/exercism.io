package clock

import "fmt"

// Clock represents the time of day in seconds
type Clock int

const day int = 60 * 24

// New initializes a Clock from hours and minutes
func New(h, m int) Clock {
	time := (h*60 + m) % day
	if time < 0 {
		time += day
	}
	return Clock(time)
}

// Add minutes to an existing Clock
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract minutes from an existing Clock
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

// String format a clock as mm:hh
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
