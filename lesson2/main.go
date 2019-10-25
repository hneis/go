package main

import (
	"fmt"

	"../lesson2/modules/functions"
)

func main() {
	// Задание 1
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

}
