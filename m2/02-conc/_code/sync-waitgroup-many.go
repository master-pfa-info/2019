package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done() // HL
		fmt.Printf("Hello from %v!\n", id)
	}

	const n = 5
	var wg sync.WaitGroup
	wg.Add(n) // HL
	for i := 0; i < n; i++ {
		go hello(&wg, i)
	}
	wg.Wait() // HL
	// STOP OMIT
}
