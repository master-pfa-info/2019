package main

import "fmt"

func main() {
	// START OMIT
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
	// STOP OMIT
}
