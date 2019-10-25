package functions

import "fmt"

func Divde3WithoutReminder(number int) (ok bool, remains int) {
	ok = false
	remains = number % 3
	if remains == 0 {
		ok = true
	}

	return
}

func Divde3WithoutReminderAndPrint(number int) {
	if ok, remains := Divde3WithoutReminder(number); ok {
		fmt.Printf("%v делится на 3 без остатка\n", number)
	} else {
		fmt.Printf("остаток от деления %v на 3 равен %v\n", number, remains)
	}

}
