package vehicle

import "fmt"

type Info struct {
	Model        string
	BuildingYear string
	TrunkVolume  float32
}

func (i Info) Print() {
	fmt.Println("General info:")
	fmt.Printf("  Model:\t\t %s\n", i.Model)
	fmt.Printf("  BuildingYear:\t\t %s\n", i.BuildingYear)
	fmt.Printf("  TrunkVolume:\t\t %f\n", i.TrunkVolume)
}

type Car struct {
	Info
	OnlyInCar        string
	TrunkVolumeUsed_ float32
	EngineRunning_   bool
	WindowOpened_    bool
}

func (c Car) Model() string {
	return c.Info.Model
}

func (c Car) BuildingYear() string {
	return c.Info.BuildingYear
}

func (c Car) TrunkVolume() float32 {
	return c.Info.TrunkVolume
}
func (c Car) TrunkVolumeUsed() float32 {
	return c.TrunkVolumeUsed_
}
func (c Car) EngineRunning() bool {
	return c.EngineRunning_
}
func (c Car) WindowOpened() bool {
	return c.WindowOpened_
}

func (c Car) VehicleDescription() {
	fmt.Println("Car description:")
	c.Info.Print()
	fmt.Println("Specific information:")
	fmt.Printf("  EngineRunning:\t %v\n", c.EngineRunning_)
	fmt.Printf("  WindowOpened:\t\t %v\n", c.WindowOpened_)
	fmt.Printf("  OnlyInCar:\t\t %v\n", c.OnlyInCar)
	fmt.Println()
}
