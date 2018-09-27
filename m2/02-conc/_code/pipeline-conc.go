package main

import "fmt"

func main() {
	// START-GEN OMIT
	gen := func(done <-chan bool, ints ...int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch) // HL
			for _, v := range ints {
				select {
				case <-done: // HL
					return
				case ch <- v:
				}
			}
		}()
		return ch
	}
	// STOP-GEN OMIT
	// START-MUL OMIT
	mul := func(done <-chan bool, ch <-chan int, a int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out) // HL
			for v := range ch {
				select {
				case <-done: // HL
					return
				case out <- a * v:
				}
			}
		}()
		return out
	}
	// STOP-MUL OMIT
	// START-ADD OMIT
	add := func(done <-chan bool, ch <-chan int, a int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out) // HL
			for v := range ch {
				select {
				case <-done: // HL
					return
				case out <- a + v:
				}
			}
		}()
		return out
	}
	// STOP-ADD OMIT
	// START-MAIN OMIT
	done := make(chan bool)
	defer close(done) // HL

	ch := gen(done, 1, 2, 3, 4)
	pipeline := mul(done, add(done, mul(done, ch, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
	// STOP-MAIN OMIT
}
