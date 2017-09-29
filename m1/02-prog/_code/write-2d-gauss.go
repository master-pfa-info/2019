package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Create("data-2d-gauss.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		g1 := 20 + 5*rand.NormFloat64()
		g2 := 10 + 20*rand.NormFloat64()
		fmt.Fprintf(f, "%v %v\n", g1, g2)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
