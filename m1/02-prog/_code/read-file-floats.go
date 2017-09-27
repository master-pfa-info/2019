package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	sum := 0.0
	for i := 0; i < 1000; i++ {
		var v float64
		_, err = fmt.Fscanf(f, "%f\n", &v)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Printf("mean= %f\n", sum/1000.0)
}
