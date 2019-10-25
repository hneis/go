package functions

import "fmt"

func IsEvenNumber(number interface{}) (result bool, err error) {
	switch v := number.(type) {
	case int:
		result = v%2 == 0
	case float32, float64:
		err = fmt.Errorf("Float numbers can't be even")
	}

	return
}
