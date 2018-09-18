package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	doWork := func(strings <-chan int) <-chan int {
		done := make(chan int)
		go func() {
			defer close(done)
			defer fmt.Println("doWork exited")
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return done
	}

	doWork(nil)
	// do some more work here.
	time.Sleep(5 * time.Second)
	fmt.Println("Done.")
	// STOP OMIT
}
