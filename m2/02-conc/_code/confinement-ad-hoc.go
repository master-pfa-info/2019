package main

import "fmt"

func main() {
	// START OMIT
	data := []int{1, 2, 3, 4} // HL

	loop := func(ch chan<- int) {
		defer close(ch)
		for _, v := range data { // HL
			ch <- v
		}
	}

	ch := make(chan int)
	go loop(ch)

	for v := range ch {
		fmt.Println(v)
	}
	// STOP OMIT
}
