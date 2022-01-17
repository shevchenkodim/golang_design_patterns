package other

import (
	"log"
	"time"
)

/*
When optimizing code, you sometimes need to quickly measure code execution time to test assumptions
instead of using profiling tools/frameworks.

In such cases, time measurements can be made using the time package and defer statements.
*/

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}

func SomeFunction() {
	defer Duration(time.Now(), "IntFactorial")

	// some logic

	// finish with auto printing spending times
}
