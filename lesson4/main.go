// Package main provides ...
package main

import (
	"fmt"
	"sort"

	"github.com/hneis/go/lesson4/chess"
	"github.com/hneis/go/lesson4/phoneBook"
	"github.com/hneis/go/lesson4/vehicle"
)

// Задание 1
func task1() {
	vehicles := []vehicle.Interface{
		vehicle.Car{
			Info: vehicle.Info{
				Model:        "Audi",
				BuildingYear: "1980",
				TrunkVolume:  200.0,
			},
			TrunkVolumeUsed_: 100.0,
			EngineRunning_:   false,
			WindowOpened_:    false,
			OnlyInCar:        "",
		},
		vehicle.NewTruckDefault("Volvo", "1950", 200),
	}

	for _, v := range vehicles {
		v.VehicleDescription()
	}
}

// Задание 2
func task2() {
	originBook := phoneBook.Contacts{
		phoneBook.Contact{
			Name:         "Dennis",
			Patronymic:   "",
			Surname:      "Rodman",
			Organization: "Chicago Bulls",
			Phones:       []int{89220301154},
		},
		phoneBook.Contact{
			Name:         "Steff",
			Patronymic:   "",
			Surname:      "Carry",
			Organization: "Lakers",
			Phones:       []int{89220301154},
		},
		phoneBook.Contact{
			Name:         "Michal",
			Patronymic:   "",
			Surname:      "Jordan",
			Organization: "Chicago Bulls",
			Phones:       []int{89220301122, 89220102233},
		},
		phoneBook.Contact{
			Name:         "Coby",
			Patronymic:   "",
			Surname:      "Braint",
			Organization: "Lakers",
			Phones:       []int{89210301111, 8123123123},
		},
		phoneBook.Contact{
			Name:         "Shaquil",
			Patronymic:   "",
			Surname:      "O'Nil",
			Organization: "Chicago Bulls",
			Phones:       []int{89420301122},
		},
	}
	bookLen := len(originBook)
	books := make(phoneBook.Contacts, bookLen, bookLen)
	copy(books, originBook)
	fmt.Println("Phone book without sorting")
	books.Print()

	copy(books, originBook)
	sort.Sort(phoneBook.ByName(books))
	fmt.Println("Sort by name")
	books.Print()

	copy(books, originBook)
	sort.Sort(phoneBook.BySurname(books))
	fmt.Println("Sort by surname")
	books.Print()

	copy(books, originBook)
	sort.Sort(phoneBook.ByOrganization(books))
	fmt.Println("Sort by organization")
	books.Print()

	copy(books, originBook)
	sort.Sort(phoneBook.ByOrganizationAndSurname(books))
	fmt.Println("Sort by organization and surname")
	books.Print()

}

// Задание 4
func task4() {
	//Конь
	knight := chess.NewKnight(1, 0, "white")
	chessboard := chess.Chessboard{
		White: [16]chess.ChessPiece{},
		Black: [16]chess.ChessPiece{},
	}
	// Слон
	bishop := chess.NewBishop(5, 0, "black")
	chessboard.SetPoint(&bishop, chess.Point{2, 6})
	fmt.Println(chessboard.CallculateAllMovies(&knight))
	chessboard.DrawValidMovie(&knight, chessboard.CallculateAllMovies(&knight))
	chessboard.DrawValidMovie(&bishop, chessboard.CallculateAllMovies(&bishop))

}

func help() {
	fmt.Println("Введите:")
	fmt.Println("  1 - чтобы посмотреть вывод Задачи №1")
	fmt.Println("  2 - чтобы посмотреть вывод Задачи №2")
	fmt.Println("  3 - чтобы посмотреть вывод Задачи №3")
	fmt.Println("  4 - чтобы посмотреть вывод Задачи №4")
	fmt.Println("  help - чтобы посмотреть вывод Задачи №4")
}

func main() {
	var action string
	help()
	for {
		if _, err := fmt.Scan(&action); err != nil {
			fmt.Printf("Something wrong: %v", err)
			return
		}

		switch action {
		case "1":
			task1()
		case "2":
			task2()
		case "3":
			fmt.Println("Интерактивный выбор задания и help и есть Задание №3")
		case "4":
			task4()
		case "help":
			help()
		case "exit":
			return
		}
	}
}
