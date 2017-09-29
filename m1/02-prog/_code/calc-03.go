package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	for _, arg := range os.Args[1:] {
		v, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Printf("sum= %d\n", sum)
}
