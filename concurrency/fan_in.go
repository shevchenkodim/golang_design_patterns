package concurrency

import "sync"

/*
Fan-In is a messaging pattern used to create a funnel for work among performers.
The message source can be clients, and the destination can be a server.
*/

// Merge Combining different channels into one channel
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Run send goroutine
	// for each incoming channel in cs.
	// send copies values from c to out
	// until c is closed, then call wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Run goroutine to close out
	// when all send goroutines are done
	// This should start after calling wg.Add.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
