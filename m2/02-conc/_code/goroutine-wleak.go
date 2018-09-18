package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// START-MAIN OMIT
	gen := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			defer fmt.Printf("gen closure exited\n") // HL
			for {
				ch <- rand.Intn(100)
			}
		}()
		return ch
	}

	ch := gen()
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i+1, <-ch)
	}
	// STOP-MAIN OMIT
}
