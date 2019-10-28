package main

import (
	"fmt"

	"../lesson2/modules/functions"
)

func main() {
	// Задание 1
	fmt.Println("Задание 1")
	if v1, err := functions.IsEvenNumber(10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v1)
	}
	if v2, err := functions.IsEvenNumber(11); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v2)
	}
	if v3, err := functions.IsEvenNumber(12.0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v3)
	}

	// Задание 2
	fmt.Println("Задание 2")
	functions.Divde3WithoutReminderAndPrint(12)
	functions.Divde3WithoutReminderAndPrint(13)
	functions.Divde3WithoutReminderAndPrint(14)

	// Задание 3
	fmt.Println("Задание 3")
	functions.Fibonachi(10)

	// Задание 4
	fmt.Println("Задание 4")
	result := functions.FindNatural(100)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}
