package main

import (
	"fmt"
	"net/http"
)

func main() {
	// START-WORK OMIT
	type Result struct {
		Err  error
		Resp *http.Response
	}
	check := func(done <-chan int, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Err: err, Resp: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	// STOP-WORK OMIT

	// START-MAIN OMIT
	done := make(chan int)
	defer close(done)

	urls := []string{"https://uca.fr", "https://badhost"}
	for res := range check(done, urls...) {
		if res.Err != nil {
			fmt.Printf("error: %v\n", res.Err)
			continue
		}
		fmt.Printf("response: %+v\n", res.Resp.Status)
	}
	// STOP-MAIN OMIT
}
