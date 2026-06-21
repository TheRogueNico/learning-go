package main

import "fmt"

func main() {
	var (
		b      byte   = 0b1111_1111
		smallI int32  = 2147483647
		bigI   uint64 = 18446744073709551615
	)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)

	// EXTRA: not part of the exercise
	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
