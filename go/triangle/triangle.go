package triangle

import (
	"math"
)

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func inBounds(a float64) bool {
	return a > 0 && !math.IsInf(a, 0)
}

func isTriangle(a, b, c float64) bool {
	return inBounds(a) && inBounds(b) && inBounds(c) && a+b >= c && b+c >= a && a+c >= b
}

func KindFromSides(a, b, c float64) Kind {
	if isTriangle(a, b, c) {
		switch {
		default:
			return Sca
		case a == b && b == c && a == c:
			return Equ
		case a == b || b == c || a == c:
			return Iso
		}
	}
	return NaT
}
