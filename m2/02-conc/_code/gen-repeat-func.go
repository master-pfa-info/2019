package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// START-REPEAT OMIT
	repeatFct := func(done <-chan bool, fct func() int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for {
				select {
				case <-done:
					return
				case ch <- fct():
				}
			}
		}()
		return ch
	}
	// STOP-REPEAT OMIT
	// START-TAKE OMIT
	take := func(done <-chan bool, ch <-chan int, n int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < n; i++ {
				select {
				case <-done:
					return
				case out <- <-ch:
				}
			}
		}()
		return out
	}
	// STOP-TAKE OMIT
	// START-MAIN OMIT

	done := make(chan bool)
	defer close(done)

	for v := range take(done, repeatFct(done, rand.Int), 10) {
		fmt.Println(v)
	}
	// STOP-MAIN OMIT
}
