package generative

func Generator(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Блокирует операцию
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}
