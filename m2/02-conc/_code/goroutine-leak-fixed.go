package main

import (
	"fmt"
	"time"
)

func main() {
	// START-WORK OMIT
	doWork := func(quit, strings <-chan int) <-chan int {
		done := make(chan int)
		go func() {
			defer close(done)
			defer fmt.Println("doWork exited")
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-quit: // HL
					fmt.Println("doWork ABORT")
					return
				}
			}
		}()
		return done
	}
	// STOP-WORK OMIT

	// START-MAIN OMIT
	quit := make(chan int)
	done := doWork(quit, nil)

	go func() { // HL
		// cancel the operation after 2 seconds
		time.Sleep(2 * time.Second)
		fmt.Println("canceling doWork goroutine...")
		close(quit)
	}()

	// do some more work here.
	time.Sleep(5 * time.Second)
	<-done // HL
	fmt.Println("Done.")
	// STOP-MAIN OMIT
}
