package main

import (
	"io"
	"os"

	"uca.fr/rot13" // HLxxx
)

func main() {
	r := rot13.NewReader(os.Stdin)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
