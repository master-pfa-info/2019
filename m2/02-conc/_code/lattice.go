package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	nn := flag.Int("xy", 2, "xy")
	flag.Parse()
	nx := *nn
	ny := *nn
	fmt.Printf("n=%d\n", compute(nx, ny))
}

// START-COMPUTE OMIT
type Point struct {
	x, y int
}

func compute(nx, ny int) int {
	var pt Point
	var wg sync.WaitGroup
	wg.Add(2)

	var ch = make(chan int)
	go func() {
		r := Point{pt.x + 1, pt.y}
		d := Point{pt.x, pt.y + 1}
		go inspect(r, &wg, nx, ny, ch)
		go inspect(d, &wg, nx, ny, ch)
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for range ch {
		sum++
	}

	return sum
}

// STOP-COMPUTE OMIT

// START-INSPECT OMIT

func inspect(pt Point, wg *sync.WaitGroup, nx, ny int, ch chan int) {
	defer wg.Done()
	if pt.x < nx {
		r := Point{pt.x + 1, pt.y}
		wg.Add(1)
		go inspect(r, wg, nx, ny, ch)
	}
	if pt.y < ny {
		d := Point{pt.x, pt.y + 1}
		wg.Add(1)
		go inspect(d, wg, nx, ny, ch)
	}

	if pt.y == ny && pt.x == nx {
		ch <- 1
	}
}

// STOP-INSPECT OMIT
