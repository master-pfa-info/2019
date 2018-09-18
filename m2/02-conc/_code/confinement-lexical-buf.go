package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	process := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buf bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buf, "%c", b)
		}
		fmt.Printf("%s\n", buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")
	go process(&wg, data[:3]) // HL
	go process(&wg, data[3:]) // HL
	wg.Wait()
	// STOP OMIT
}
