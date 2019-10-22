// Package exchanger provides money exchanger
package exchanger

import "fmt"

func RUBToUSD(rate float64) {
	var value float64

	fmt.Println("Конвертер валют RUB => $")
	fmt.Print("Введите RUB: ")

	fmt.Scan(&value)

	result := value * rate

	fmt.Printf("%0.2f рублей равна %0.2f в долларах\n", value, result)
}
