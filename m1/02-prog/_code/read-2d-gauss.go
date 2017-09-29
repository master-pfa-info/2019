package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, err := os.Open("data-2d-gauss.txt")
	if err != nil {
		panic(err)
	}

	var g1, g2 []float64
	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		var v1, v2 float64
		_, err = fmt.Sscanf(scan.Text(), "%f %f", &v1, &v2)
		if err != nil {
			panic(err)
		}
		g1 = append(g1, v1)
		g2 = append(g2, v2)
	}

	if err = scan.Err(); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}

	fmt.Printf("g1: mean=%v stddev=%v\n", mean(g1), stddev(g1))
	fmt.Printf("g2: mean=%v stddev=%v\n", mean(g2), stddev(g2))
}

// mean computes the un-weighted mean of the data set.
//  sum_i {w_i * x_i} / sum_i {w_i} . where all w_i are = 1.
func mean(x []float64) float64 {
	sum := 0.0
	for i := range x {
		sum += x[i]
	}
	return sum / float64(len(x))
}

// stddev computes the sqrt(variance), where the variance is
//  \sum_i w_i (x_i - mean)^2 / (sum_i w_i - 1) . where all w_i are 1.
func stddev(x []float64) float64 {
	m := mean(x)
	s := 0.0
	for i := range x {
		v := x[i]
		d := v - m
		s += d * d
	}
	variance := s / float64(len(x)-1)
	return math.Sqrt(variance)
}
