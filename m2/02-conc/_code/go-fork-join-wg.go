package main

import "sync"

// START OMIT
func main() {
	println("in main")
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		println("hello there!")
	}
	wg.Add(1)
	go sayHello() // fork // HL
	wg.Wait()     // join // HL
	println("bye.")
}

// STOP OMIT
