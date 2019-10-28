package functions

import "fmt"

func Fibonachi(n int) {
	gen := genFib()
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", gen())
	}
	fmt.Printf("\n")
}

func genFib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}
