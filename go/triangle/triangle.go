package triangle

import "math"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func isValid(n float64) bool {
	return !(math.IsNaN(n) || math.IsInf(n, 0) || n <= 0)
}

func KindFromSides(a, b, c float64) Kind {
	if !(isValid(a) && isValid(b) && isValid(c)) {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
