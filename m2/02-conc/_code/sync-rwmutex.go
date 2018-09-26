package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	// START-HEAD OMIT
	type locker interface {
		Lock()
		Unlock()
	}

	producer := func(wg *sync.WaitGroup, l locker) {
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
		wg.Done()
	}

	observer := func(wg *sync.WaitGroup, l locker) {
		l.Lock()
		l.Unlock()
		wg.Done()
	}
	// STOP-HEAD OMIT

	// START-MAIN OMIT
	test := func(count int, mux, rwmux locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beg := time.Now()
		go producer(&wg, mux)
		for i := count; i > 0; i-- {
			go observer(&wg, rwmux)
		}
		wg.Wait()
		return time.Since(beg)
	}
	// STOP-MAIN OMIT

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw, "%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}
