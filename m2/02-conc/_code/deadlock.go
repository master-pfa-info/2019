package main

import (
	"fmt"
	"sync"
	"time"
)

// START-VALUE OMIT
type value struct {
	mu    sync.Mutex
	value int
}

// STOP-VALUE OMIT

func main() {
	// START-MAIN OMIT
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         // HL
		defer v1.mu.Unlock() // HL

		time.Sleep(2 * time.Second)
		v2.mu.Lock()         // HL
		defer v2.mu.Unlock() // HL

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
	// STOP-MAIN OMIT
}
