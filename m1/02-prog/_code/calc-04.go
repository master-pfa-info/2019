package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flagVerbose := flag.Bool("v", false, "enable/disable verbose mode") // HLflag
	flag.Parse()                                                        // HLflag

	sum := 0
	if *flagVerbose { // HLflag
		fmt.Printf("%v\n", sum)
	}
	for _, arg := range flag.Args() { // HLflag
		v, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		sum += v

		if *flagVerbose { // HLflag
			fmt.Printf("+ %v -> %v\n", v, sum)
		}
	}
	if *flagVerbose { // HLflag
		fmt.Printf("===============\n")
	}
	fmt.Printf("sum= %v\n", sum)
}
