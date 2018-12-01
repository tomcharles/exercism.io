// Package clock does
package clock

import "fmt"

// Clock says
type Clock struct {
	hour   int
	minute int
}

// New says
func New(h, m int) Clock {
	addHours := m / 60
	hour := 0 + (h+addHours)%24
	if hour < 0 {
		hour = 24 + hour
	}
	if m < 0 && addHours != 0 {
		hour = hour - 1
	}

	minute := 0 + (m % 60)
	if minute < 0 {
		minute = 60 + minute
	}

	return Clock{hour: hour, minute: minute}
}

// Add says
func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

// Subtract syas
func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

// String says
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
