package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/master-pfa-info/mcpi"
)

// START OMIT

func main() {
	n := flag.Int("n", 1e7, "MC sample size")
	flag.Parse()

	mcpi.Wait() // HL
	fmt.Printf("pi(%d) = %v\n", *n, pi(*n))
	mcpi.Quit() // HL
}

func pi(n int) float64 {
	inside := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
		mcpi.Plot(x, y) // HL
	}
	return 4 * float64(inside) / float64(n)
}

// STOP OMIT
