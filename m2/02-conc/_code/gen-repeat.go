package main

import "fmt"

func main() {
	// START-REPEAT OMIT
	repeat := func(done <-chan bool, values ...int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case ch <- v:
					}
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

	for v := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Println(v)
	}
	// STOP-MAIN OMIT
}
