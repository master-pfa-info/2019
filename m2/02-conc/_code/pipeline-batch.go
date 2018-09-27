package main

import "fmt"

func main() {
	// START-MUL OMIT
	mul := func(vs []int, a int) []int {
		out := make([]int, len(vs))
		for i, v := range vs {
			out[i] = a * v
		}
		return out
	}
	// STOP-MUL OMIT

	// START-ADD OMIT
	add := func(vs []int, a int) []int {
		out := make([]int, len(vs))
		for i, v := range vs {
			out[i] = a + v
		}
		return out
	}
	// STOP-ADD OMIT

	// START-MAIN OMIT
	ints := []int{1, 2, 3, 4}
	for _, v := range mul(add(mul(ints, 2), 1), 2) {
		fmt.Println(v)
	}
	// STOP-MAIN OMIT
}
