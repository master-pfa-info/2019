package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	c := sync.NewCond(&sync.Mutex{})
	q := make([]int, 0, 10)

	remove := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock() // HL
		v := q[0]
		q = q[1:]
		fmt.Println("removed from queue:", v)
		c.L.Unlock() // HL
		c.Signal()   // signals 1 goroutine // HL
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // HL
		for len(q) == 2 {
			c.Wait() // HL
		}
		fmt.Println("adding to queue:", i)
		q = append(q, i)
		go remove(1 * time.Second)
		c.L.Unlock() // HL
	}
	// STOP OMIT
}
