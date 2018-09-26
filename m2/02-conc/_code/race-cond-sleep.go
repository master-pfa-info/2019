package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var data int
	go func() {
		data++
	}()
	time.Sleep(2 * time.Second) // bad! // HL
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
	// STOP OMIT
}
