package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// define the command line flags
	verbose := flag.Bool("v", false, "enable verbose mode")
	ptMin := flag.Float64("pt-min", 25.0, "min Pt cut for H->bb analysis")
	btag := flag.Int("btag", 2, "number of b-tag for H->bb analysis")

	// parse the command line arguments of the program and assign values to flags.
	flag.Parse()

	fmt.Printf("verbose=   %v\n", *verbose)
	fmt.Printf("pt-min=    %f\n", *ptMin)
	fmt.Printf("b-tag=     %d\n", *btag)
	fmt.Printf("os.Args=   %v\n", os.Args)
	fmt.Printf("flag.Args= %v\n", flag.Args())
}
