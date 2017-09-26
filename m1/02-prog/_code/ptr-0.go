package main

import "fmt"

func main() {
	var i = 11 // an integer i, with value 11
	fmt.Printf("i=%d, addr=%#x\n", i, &i)

	var p *int = &i // a pointer to the integer i
	fmt.Printf("p=%#x, addr=%#x\n", p, &p)

	*p += 22 // dereference p, and assign a new value

	fmt.Printf("i=%d, addr=%#x\n", i, &i)
	fmt.Printf("p=%#x, addr=%#x\n", p, &p)

	var pp *float64
	fmt.Printf("pp=%v %#x\n", pp, pp)
}
