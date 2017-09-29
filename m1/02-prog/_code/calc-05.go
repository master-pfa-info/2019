package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	flagVerbose := flag.Bool("v", false, "enable/disable verbose mode")
	flagFile := flag.String("f", "", "path to a file holding values")
	flagStats := flag.Bool("stats", false, "enable computation+display of stats")
	flag.Parse()

	var values []float64
	if *flagFile != "" {
		values = readFile(*flagFile)
	} else {
		values = readArgs()
	}
	run(values, *flagVerbose, *flagStats)
}

func run(vs []float64, verbose, stats bool) {
	sum := 0.0
	if verbose {
		fmt.Printf("%v\n", sum)
	}
	for _, v := range vs {
		sum += v
		if verbose {
			fmt.Printf("+ %v -> %v\n", v, sum)
		}
	}
	if verbose {
		fmt.Printf("===============\n")
	}
	fmt.Printf("sum= %v\n", sum)
	if stats {
		fmt.Printf("mean=   %v\n", mean(vs))
		fmt.Printf("stddev= %v\n", stddev(vs))
	}
}

func readFile(fname string) []float64 {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	var o []float64
	scan := bufio.NewScanner(f) // need to import "bufio"
	for scan.Scan() {
		var v float64
		txt := scan.Text()
		_, err := fmt.Sscanf(txt, "%f", &v)
		if err != nil {
			panic(err)
		}
		o = append(o, v)
	}
	if err = scan.Err(); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
	return o
}

func readArgs() []float64 {
	var o []float64
	for _, arg := range flag.Args() {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		o = append(o, v)
	}
	return o
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
