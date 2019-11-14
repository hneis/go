package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func task1() {
	cancelCh := make(chan struct{})
	go spinner(cancelCh, 50*time.Millisecond)

	// например так
	whait := time.After(time.Second * 2)
	select {
	case <-whait:
		close(cancelCh)
	}
	// можно так
	// fmt.Scanln()
}

func spinner(cancelCh <-chan struct{}, delay time.Duration) {
	for {
		select {
		case <-cancelCh:
			return
		default:
			for _, r := range "-\\|/" {
				fmt.Printf("%c\r", r)
				time.Sleep(delay)
			}
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func task2() {
	cancelChanel := make(chan struct{})
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func(n int, cc chan<- struct{}) {
		for x := 0; x < n; x++ {
			naturals <- x
		}
		close(cc)
	}(100000, cancelChanel)

	// возведение в квадрат
	go func(cc <-chan struct{}) {
		for {
			select {
			case x := <-naturals:
				squares <- x * x
			case <-cc:
				close(squares)
				return
			}
		}
	}(cancelChanel)

	// печать
	var str string
	wString := bytes.NewBufferString(str)
	bufio.NewWriter(wString)
	for sq := range squares {
		wString.WriteString(strconv.Itoa(sq))
		wString.WriteString("\n")
	}
	fmt.Println(wString)
}

func handleConn(ctx Context, c net.Conn) {
	defer c.Close()
	defer ctx.wg.Done()
	for {
		select {
		case <-ctx.close:
			fmt.Println("Close connection from ", c.RemoteAddr())
			return
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

type Context struct {
	wg    *sync.WaitGroup
	close chan struct{}
}

func task3() {
	context := Context{
		wg:    &sync.WaitGroup{},
		close: make(chan struct{}),
	}
	go func(ctx Context) {
		listener, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			fmt.Println("Accept connection from ", conn.RemoteAddr())
			ctx.wg.Add(1)
			// Запускаем как goroutine(возможность нескольких подключений к серверу)
			go handleConn(ctx, conn)
		}
	}(context)
	for {
		var input string
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Server will be close")
			close(context.close)
			break
		}
	}
	context.wg.Wait()
	fmt.Println("All connection closed\nExit")
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	urls := []string{
		"duckduckgo.com",
		"www.google.com",
		"www.rambler.com",
		"www.ya.ru",
	}
	for _, url := range urls {
		go func(url string, r chan<- string) {
			r <- request(url)
		}(url, responses)
	}

	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	conn, err := net.Dial("tcp", hostname+":80")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%s time elapsed %v", hostname, time.Since(start))
	return hostname
}

func task4() {
	fmt.Println("Find fast host:")
	fastUrl := mirroredQuery()

	// Ждем 1 секунду, после чего выводим самого быстрого
	timer := time.After(time.Second)
	select {
	case <-timer:
		fmt.Printf("The fastest is %s\n", fastUrl)
	}
}

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
	// task1()
	// task2()
	// task3()
	// task4()
	task5()
}
