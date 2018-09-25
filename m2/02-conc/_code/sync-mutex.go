package main

import (
	"fmt"
	"sync"
)

func main() {
	// START-DECL OMIT
	var count = 0
	var lock sync.Mutex

	incr := func() {
		lock.Lock()         // HL
		defer lock.Unlock() // HL
		count++
		fmt.Printf("incrementing: %d\n", count)
	}

	decr := func() {
		lock.Lock()         // HL
		defer lock.Unlock() // HL
		count--
		fmt.Printf("decrementing: %d\n", count)
	}
	// STOP-DECL OMIT
	// START-MAIN OMIT
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decr()
		}()
	}

	wg.Wait()
	fmt.Printf("arithmetic complete. count=%d\n", count)
	// STOP-MAIN OMIT
}
