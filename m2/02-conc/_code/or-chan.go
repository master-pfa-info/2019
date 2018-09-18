package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	sig := func(after time.Duration) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			time.Sleep(after)
		}()
		return ch
	}

	fmt.Printf("starting or-chan...\n")
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
	// STOP OMIT

	time.Sleep(2 * time.Second)
}

func or(chans ...<-chan int) <-chan int {
	switch len(chans) {
	case 0: return nil
	case 1: return chans[0]
	}
	done := make(chan int)
	go func() {
		defer close(done)                // HL
		defer println("goroutine done.") // OMIT
		switch len(chans) {
		case 2:
			select {
			case <-chans[0]:
			case <-chans[1]:
			}
		default:
			select {
			case <-chans[0]:
			case <-chans[1]:
			case <-chans[2]:
			case <-or(append(chans[3:], done)...): // HL
			}
		}
	}()
	return done
}
