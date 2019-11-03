// Package vehicle provides ...
package vehicle

type Interface interface {
	Model() string
	BuildingYear() string
	TrunkVolume() float32
	VehicleDescription()
	TrunkVolumeUsed() float32
	EngineRunning() bool
	WindowOpened() bool
}
