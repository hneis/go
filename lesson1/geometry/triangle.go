// Package geometry provides ...
package geometry

import "math"

func TriangleArea(a, b float64) float64 {
	return 0.5 * a * b
}

func TrianglePerimeter(a, b, c float64) float64 {
	return a + b + c
}

func TriangleHypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
