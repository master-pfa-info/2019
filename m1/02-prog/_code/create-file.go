package main

import (
	"os"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}

	f.WriteString("hello world\n0 1 2 3 4 5\ngood bye\n")

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
