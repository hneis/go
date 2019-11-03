// Package vehicle provides ...
package vehicle

import "fmt"

type Engine struct {
	IsRuning bool
	Model    string
	Volume   float32
}

type Truck struct {
	TruckModel       string
	TruckBuild       string
	TruckTrunkVolume float32
	Used             float32
	Engine           Engine
}

func NewTruckDefault(model string, year string, trunkVolume float32) Truck {
	return Truck{
		TruckModel:       model,
		TruckBuild:       year,
		TruckTrunkVolume: trunkVolume,
		Used:             100.0,
		Engine: Engine{
			IsRuning: true,
			Model:    "V8",
			Volume:   12.0,
		},
	}
}

func (t Truck) Model() string {
	return t.TruckModel
}

func (t Truck) BuildingYear() string {
	return t.TruckBuild
}

func (t Truck) TrunkVolume() float32 {
	return t.TruckTrunkVolume
}

func (t Truck) VehicleDescription() {
	fmt.Println("Truck description:")
	fmt.Println("General info:")
	fmt.Printf("  Model:\t\t %s\n", t.Model())
	fmt.Printf("  BuildingYear:\t\t %s\n", t.BuildingYear())
	fmt.Printf("  TrunkVolume:\t\t %f\n", t.TrunkVolume())
	fmt.Println(" Engine specification:")
	fmt.Printf("  Model:\t\t %v\n", t.Engine.Model)
	fmt.Printf("  Is runing:\t\t %v\n", t.Engine.IsRuning)
	fmt.Printf("  Volume:\t\t %v\n", t.Engine.Volume)
	fmt.Println()
}

func (t Truck) TrunkVolumeUsed() float32 {
	return t.Used
}

func (t Truck) EngineRunning() bool {
	return t.Engine.IsRuning
}

func (t Truck) WindowOpened() bool {
	return false
}
