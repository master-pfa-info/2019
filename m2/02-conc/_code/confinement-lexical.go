package main

import "fmt"

func main() {
	// START OMIT
	producer := func() <-chan int {
		results := make(chan int) // HL
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(ch <-chan int) { // HL
		for v := range ch {
			fmt.Printf("received: %d\n", v)
		}
		fmt.Printf("done receiving values\n")
	}

	out := producer() // 'out' can only be read from.
	consumer(out)
	// STOP OMIT
}
