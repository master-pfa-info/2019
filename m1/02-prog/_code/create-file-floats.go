package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(f, "%v\n", rand.NormFloat64())
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
