package main

import (
	"flag"
	"fmt"

	"github.com/master-pfa-info/mcpi"
)

func main() {
	n := flag.Int("n", 1e7, "MC sample size")
	flag.Parse()

	mcpi.Wait()
	fmt.Printf("pi(%d) = %v\n", *n, pi(*n))
	mcpi.Quit()
}

func pi(n int) float64 {
	// ???
}
