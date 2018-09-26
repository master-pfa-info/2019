package main

import "sync"

func main() {
	// START OMIT
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock() // HL
	for condition() == false {
		c.Wait() // HL
	}
	c.L.Unlock() // HL
	// STOP OMIT
}

func condition() bool {
	return false
}
