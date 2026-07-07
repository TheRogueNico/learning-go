package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Double[T Number](n T) T {
	return n * 2
}

func main() {
	var unum uint = 16
	num := 8
	fnum := 4.5

	fmt.Println(Double(num))
	fmt.Println(Double(unum))
	fmt.Println(Double(fnum))
}
