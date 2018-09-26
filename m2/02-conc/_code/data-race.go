package main

import "fmt"

func main() {
	// START OMIT
	var data int
	go func() { data++ }() // HL
	if data == 0 {         // HL
		fmt.Printf("the value is 0. !!\n")
	} else {
		fmt.Printf("the value is %d\n", data) // HL
	}
	// STOP OMIT
}
