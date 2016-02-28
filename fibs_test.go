package fibs

import (
	"fmt"
)

func ExampleTake() {
	ns := New(1, 1).Take(10)
	fmt.Print(ns)
	// Output:
	// [1 1 2 3 5 8 13 21 34 55]
}
