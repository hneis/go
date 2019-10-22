package main

import (
	"fmt"
)

func main() {
	fmt.Println("Рассчитать сумму вклада через 5 лет")

	var deposit, rate float64

	fmt.Printf("Сумма вклада: ")
	fmt.Scan(&deposit)
	fmt.Printf("Годовой процент: ")
	fmt.Scan(&rate)
	fmt.Printf("Сумма вклада через пять лет составит: %0.3f\n", deposit+(deposit*rate*5))
}
