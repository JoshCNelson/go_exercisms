// Triangle package prvoides a function for determining whether 3 sides form a triangle and if so what kind
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides takes 3 sides and determines the Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isNotTriangle(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a != b && a != c && b != c:
		return Sca
	default:
		return Iso
	}
}

func isNotTriangle(a, b, c float64) bool {
	return hasNonExistantSide(a, b, c) ||
		hasSidesNotGreaterThanThird(a, b, c) ||
		hasInfSide(a, b, c) ||
		hasNaNSide(a, b, c)
}

func hasSidesNotGreaterThanThird(a, b, c float64) bool {
	return (a+b) < c || (a+c) < b || (b+c) < a
}

func hasNonExistantSide(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0
}

func hasNaNSide(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
}

func hasInfSide(a, b, c float64) bool {
	return math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)
}
