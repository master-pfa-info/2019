package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	sum := 0.0
	n := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		var v float64
		var txt = scan.Text()
		_, err = fmt.Sscanf(txt, "%f", &v)
		if err != nil {
			panic(err)
		}
		sum += v
		n++
	}
	fmt.Printf("mean= %f\n", sum/float64(n))
}
