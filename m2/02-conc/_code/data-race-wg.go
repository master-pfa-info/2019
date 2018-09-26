package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var mux sync.Mutex
	var data int
	go func() {
		mux.Lock() // HL
		data++
		mux.Unlock() // HL
	}()
	mux.Lock() // HL
	if data == 0 {
		fmt.Printf("the value is 0. !!\n")
	} else {
		fmt.Printf("the value is %d\n", data)
	}
	mux.Unlock() // HL
	// STOP OMIT
}
