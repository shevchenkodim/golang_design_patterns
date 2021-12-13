package structural

import "log"

type Object func(int) int

func logDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Start => ", n)

		result := fn(n)

		log.Println("End => ", result)

		return result
	}
}

func double(n int) int {
	return n * 2
}

func SomeDecorator() {
	f := logDecorate(double)
	f(5)
}
