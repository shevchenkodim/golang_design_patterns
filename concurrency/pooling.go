package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

/*
Pooling Pattern
The main idea behind Pooling pattern is to have:
 - a channel that provides a signaling semantics. Unbuffered channel is used to have a guarantee a goroutine has received a signal
 - multiple goroutines that pool that channel for work
 - a goroutine that sends work via channel
*/

func DoConcurrencyPooling() {
	ch := make(chan string)

	g := runtime.NumCPU()
	for e := 0; e < g; e++ {
		// a new goroutine is created for each employee
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : received signal : %s\n", emp, p)
			}

			fmt.Printf("employee %d : revieved shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "do it"

		fmt.Println("manager : sent signal :", w)
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
}
