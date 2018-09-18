package main

import (
	"fmt"
	"net/http"
)

func main() {
	// START-WORK OMIT
	check := func(done <-chan int, urls ...string) <-chan *http.Response {
		resps := make(chan *http.Response)
		go func() {
			defer close(resps)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("error: %v\n", err) // HL
					continue
				}
				select {
				case <-done:
					return
				case resps <- resp:
				}
			}
		}()
		return resps
	}
	// STOP-WORK OMIT

	// START-MAIN OMIT
	done := make(chan int)
	defer close(done)

	urls := []string{"https://uca.fr", "https://badhost"}
	for resp := range check(done, urls...) {
		fmt.Printf("response: %+v\n", resp.Status)
	}
	// STOP-MAIN OMIT
}
