package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := make([]int, 100)

	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	for _, v := range numbers {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%3 == 0:
			fmt.Println("Three!")
		case v%2 == 0:
			fmt.Println("Two!")
		default:
			fmt.Println("Never mind.")
		}
	}
}
