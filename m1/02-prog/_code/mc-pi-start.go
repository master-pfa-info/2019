package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 1e7, "MC sample size")
	flag.Parse()
	fmt.Printf("pi(%d) = %v\n", *n, pi(*n))
}

func pi(n int) float64 {
	// ???
}
