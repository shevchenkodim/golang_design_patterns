package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("NumCPU => ", runtime.NumCPU())

	// decorator
	//structural.SomeDecorator()

	// Generator
	//generative.Generator(1, 100)

	// Singleton
	//generative.CheckSingleton()
}
