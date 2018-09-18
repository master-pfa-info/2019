package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// START-MAIN OMIT
	gen := func(done <-chan int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			defer fmt.Printf("gen closure exited\n")
			for {
				select {
				case ch <- rand.Intn(100): // HL
				case <-done: // HL
					return
				}
			}
		}()
		return ch
	}

	done := make(chan int)
	ch := gen(done)
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i+1, <-ch)
	}
	close(done)
	// STOP-MAIN OMIT

	// simulate some more work...
	time.Sleep(2 * time.Second)
}
