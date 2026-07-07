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

	fmt.Println(numbers)
}
