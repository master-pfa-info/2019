package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("args[%d] = %q (%T)\n", i, arg, arg)
	}
}
