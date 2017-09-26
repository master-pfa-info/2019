package main

import "fmt"

func main() {
	var i = 42
	fmt.Printf("i=%d\n", i)

	add_1(i, 10)
	fmt.Printf("i=%d\n", i)

	j := add_2(i, 10)
	fmt.Printf("i=%d\n", i)
	fmt.Printf("j=%d\n", j)

	add_3(&i, 22)
	fmt.Printf("i=%d\n", i)
}

func add_1(i, n int)      { i += n }
func add_2(i, n int) int  { i += n; return i }
func add_3(i *int, n int) { *i += n }
