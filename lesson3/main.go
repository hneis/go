// Package main provides
package main

import (
	"fmt"

	"github.com/hneis/go/lesson3/queue"
)

type Info struct {
	Model        string
	BuildingYear string
	TrunkVolume  float32
}

func (i Info) Print() {
	fmt.Println("General info:")
	fmt.Printf("  Model: %s\n", i.Model)
	fmt.Printf("  BuildingYear: %s\n", i.BuildingYear)
	fmt.Printf("  TrunkVolume: %f\n", i.TrunkVolume)
}

type Car struct {
	Info
	TrunkVolumeUsed float32
	EngineRunning   bool
	WindowOpened    bool
	OnlyInCar       string
}

func (c Car) Print() {
	fmt.Println("Car description:")
	c.Info.Print()
	fmt.Println("Specific information:")
	fmt.Printf("  EngineRunning: %v\n", c.EngineRunning)
	fmt.Printf("  WindowOpened: %v\n", c.WindowOpened)
	fmt.Printf("  OnlyInCar: %v\n", c.OnlyInCar)
	fmt.Println()
}

type Truck struct {
	Info
	TrunkVolumeUsed float32
	EngineRunning   bool
	WindowOpened    bool
	OnlyInTruck     string
}

func main() {
	cars := []interface{}{
		Car{
			Info: Info{
				Model:        "Audi",
				BuildingYear: "1993",
				TrunkVolume:  200,
			},
			TrunkVolumeUsed: 150.0,
			EngineRunning:   false,
			WindowOpened:    true,
			OnlyInCar:       "fast speed",
		},
		Car{
			Info: Info{
				Model:        "BMW",
				BuildingYear: "1980",
				TrunkVolume:  192.0,
			},
			TrunkVolumeUsed: 100.0,
			EngineRunning:   true,
			WindowOpened:    false,
			OnlyInCar:       "very fast speed",
		},
		Truck{
			Info: Info{
				Model:        "Volvo",
				BuildingYear: "1950",
				TrunkVolume:  2000.0,
			},
			TrunkVolumeUsed: 100.0,
			EngineRunning:   false,
			WindowOpened:    false,
			OnlyInTruck:     "Big trunk",
		},
	}

	for _, car := range cars {
		switch v := car.(type) {
		case Car:
			v.Print()
		case Truck:
			v.Info.Print()
			fmt.Println(v.OnlyInTruck)

		}
	}

	// Задание 3.
	q := queue.NewQueue(100)
	q.Push(100)
	q.Push(101)
	q.Push(102)
	q.Push(103)
	q.Push(104)
	q.Print()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Print()
}
