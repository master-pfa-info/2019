package main

import (
	"fmt"

	"uca.fr/ints"
)

func main() {
	slice := ints.Gen(20)
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("sum = %v\n", ints.Sum(slice))
}
