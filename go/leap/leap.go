// Package leap can tell if a year is a leap year
package leap

// IsLeapYear returns a boolean describing if the given year is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
