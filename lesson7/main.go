package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Car struct {
	speed int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano())
	return &Car{
		speed: rand.Intn(121),
	}
}

// func (c Car) Path(params) {

// }

type Track struct {
	cars     []*Car
	trackLen float64
	position map[*Car]float64
	wg       *sync.WaitGroup
	mutex    *sync.Mutex
	finished map[*Car]bool
}

func (t Track) drawTrack(duration time.Duration) string {
	const TRAK_LEN_DRAW = 80
	per := TRAK_LEN_DRAW / t.trackLen

	track := fmt.Sprintf("├%s┤", strings.Repeat("─", TRAK_LEN_DRAW-2))
	// for i := 0; i < 100; i++ {
	for _, car := range t.cars {
		carPositionOnTrack := float64(car.speed) * duration.Hours()
		carDrawPosition := int(float64(carPositionOnTrack) * per)
		before := carDrawPosition - 1
		after := TRAK_LEN_DRAW - carDrawPosition
		path := fmt.Sprintf("%s%s%s",
			strings.Repeat(" ", before),
			"▷",
			strings.Repeat(" ", after))
		_ = path
	}

	// fmt.Sprintln(strings.Repeat("  ", in), "█",
	// 	strings.Repeat("  ", gortineCount-in),
	// 	"Thread", in,
	// 	"Iteration", j, strings.Repeat("<", j))
	return track
}

func (t *Track) Start() {
	for _, car := range t.cars {
		t.wg.Add(1)
		t.mutex.Lock()
		t.finished[car] = false
		t.mutex.Unlock()

		go func(track *Track, car *Car) {
			rand.Seed(time.Now().UnixNano())
			defer track.wg.Done()
			for {
				track.mutex.Lock()

				if track.finished[car] {
					track.mutex.Unlock()
					return
				}
				track.mutex.Unlock()
				timer := time.After(time.Second)
				select {
				case <-timer:
					track.mutex.Lock()

					axil := rand.Float64()
					track.position[car] += float64(car.speed) * axil * time.Duration(time.Minute*10).Hours()
					if t.position[car] >= t.trackLen {
						t.position[car] = t.trackLen
						t.finished[car] = true
					}
					track.mutex.Unlock()
				}
			}
		}(t, car)
	}

	// TODO Вот эту горутину надо бы завершать
	go func(track *Track) {
		const TRAK_LEN_DRAW = 80
		per := TRAK_LEN_DRAW / track.trackLen
		drawingTrack := fmt.Sprintf("├%s┤", strings.Repeat("─", TRAK_LEN_DRAW-2))
		fmt.Println(drawingTrack)
		for {
			timer := time.After(time.Millisecond * 50)
			select {
			case <-timer:
				track.mutex.Lock()
				var trackStr string
				buffer := bytes.NewBufferString(trackStr)
				writer := bufio.NewWriter(buffer)
				for _, car := range track.cars {
					carPositionOnTrack := track.position[car]
					carDrawPosition := int(float64(carPositionOnTrack) * per)
					before := 0
					if carDrawPosition > 0 {
						before = carDrawPosition - 1
					}
					after := TRAK_LEN_DRAW - carDrawPosition
					if after < 0 {
						after = TRAK_LEN_DRAW
					}
					path := fmt.Sprintf("%s%s%s%f\n",
						strings.Repeat(" ", before),
						"▷",
						strings.Repeat(" ", after),
						track.position[car])
					writer.WriteString(path)
				}
				writer.Flush()
				fmt.Print("\033[H\033[2J")
				fmt.Println(drawingTrack)
				fmt.Printf("%s\r", buffer)
				track.mutex.Unlock()
			}
		}
	}(t)
}

func (t *Track) AddCar(car *Car) {
	t.cars = append(t.cars, car)
	t.position[car] = 0
	t.finished[car] = false
}

func task5() {
	track := Track{
		trackLen: 160,
		wg:       &sync.WaitGroup{},
		mutex:    &sync.Mutex{},
		position: map[*Car]float64{},
		finished: map[*Car]bool{},
	}
	track.AddCar(NewCar())
	track.AddCar(NewCar())
	track.AddCar(NewCar())
	track.AddCar(NewCar())
	track.AddCar(NewCar())
	track.Start()
	track.wg.Wait()
}

func main() {
	task5()
}
