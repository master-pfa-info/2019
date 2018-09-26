package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var count int
	incr := func() { count++ }
	var once sync.Once

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(incr)
		}()
	}
	wg.Wait()
	fmt.Printf("count= %v\n", count)
	// STOP OMIT
}
