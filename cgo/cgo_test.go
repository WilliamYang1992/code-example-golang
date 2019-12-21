package cgo

import "fmt"

func Example() {
	sum := Sum(1, 2)
	fmt.Println(sum)
	// Output:
	// 3
}
