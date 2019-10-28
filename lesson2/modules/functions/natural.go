package functions

import "fmt"

func FindNatural(N int) (result []int) {
	first := 2
	end := 50
	for len(result) < N {
		var tmp []int
		tmp, first, end = genNatural(first, end)
		result = append(result, tmp...)
		p := 2
		for i := 0; i < len(result); i++ {
			for v := 2; v < result[len(result)-1]; v++ {
				remove := v * p
				index, err := findIndex(result, remove)
				if err == nil {
					result = append(result[:index], result[index+1:]...)
				}
			}
			p = findNextP(result, p)
		}
	}
	result = append(result[:100])

	return result
}

func genNatural(first, end int) (result []int, newFirst, newEnd int) {
	for i := first; i <= end; i++ {
		result = append(result, i)
	}
	newFirst = end + 1
	newEnd += newFirst + 100

	return
}

func findIndex(array []int, remove int) (index int, err error) {
	for index = 0; index < len(array); index++ {
		if array[index] == remove {
			return
		}
	}
	return -1, fmt.Errorf("not found")
}

func findNextP(array []int, oldP int) (result int) {
	for i := 0; i < len(array); i++ {
		if array[i] > oldP {
			result = array[i]
			break
		}
	}
	return
}
