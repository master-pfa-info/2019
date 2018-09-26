package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var count int
	incr := func() { count++ }
	decr := func() { count-- }

	var once sync.Once
	once.Do(incr)
	once.Do(decr)

	fmt.Printf("count = %d\n", count)
	// STOP OMIT
}
