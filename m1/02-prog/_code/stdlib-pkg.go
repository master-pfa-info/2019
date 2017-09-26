package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 10; i++ {
		f := float64(i)
		fmt.Printf("sqrt(%v) = %f\n", f, math.Sqrt(f))
	}
	fmt.Printf("pi = %v\n", math.Pi)
}
