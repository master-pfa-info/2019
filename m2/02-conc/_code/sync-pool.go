package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new value")
			return 1
		},
	}
	_ = pool.Get()
	v := pool.Get()
	pool.Put(v)
	_ = pool.Get()
	// STOP OMIT
}
