package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fct func()) {
		var running sync.WaitGroup
		running.Add(1)
		go func() {
			running.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fct()
		}()
		running.Wait()
	}
	// STOP OMIT
	// START-MAIN OMIT
	var click sync.WaitGroup
	click.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("maximizing window")
		click.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("displaying annoying dialog box")
		click.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("mouse clicked")
		click.Done()
	})

	button.Clicked.Broadcast()
	click.Wait()
	// STOP-MAIN OMIT
}
