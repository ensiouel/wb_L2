package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("exit after %v\n", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(1)

	wait := func(wg *sync.WaitGroup, c <-chan interface{}) {
		for v := range c {
			out <- v
		}

		wg.Done()
	}

	for _, channel := range channels {
		go wait(&wg, channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
