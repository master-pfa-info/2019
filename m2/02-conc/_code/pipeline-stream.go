package main

import "fmt"

func main() {
	// START-MUL OMIT
	mul := func(v, a int) int {
		return v * a
	}
	// STOP-MUL OMIT

	// START-ADD OMIT
	add := func(v, a int) int {
		return a + v
	}
	// STOP-ADD OMIT

	// START-MAIN OMIT
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(mul(add(mul(v, 2), 1), 2))
	}
	// STOP-MAIN OMIT
}
