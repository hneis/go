package main

import (
	"fmt"

	"../lesson1/geometry"
)

const (
	DOLLAR_EXCHANGE = 65
)

func main() {
	fmt.Println("Площадь, периметр и гипотенуза прямоугольного треугольника")

	var a, b float64

	fmt.Printf("Введите сторону a: ")
	fmt.Scan(&a)
	fmt.Printf("Введите сторону b: ")
	fmt.Scan(&b)
	c := geometry.TriangleHypotenuse(a, b)
	fmt.Printf("Площадь S=%0.3f\n", geometry.TriangleArea(a, b))
	fmt.Printf("Периметр P=%0.3f\n", geometry.TrianglePerimeter(a, b, c))
	fmt.Printf("Гипотенуза c=%0.3f\n", c)
}
