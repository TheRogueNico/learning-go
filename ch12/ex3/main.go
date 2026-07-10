package main

import (
	"fmt"
	"math"
	"sync"
)

func buildSqrtMap() map[int]float64 {
	m := make(map[int]float64, 100_000)
	for i := range 100_000 {
		m[i] = math.Sqrt(float64(i))
	}

	return m
}

var getSqrtMap = sync.OnceValue(buildSqrtMap)

func main() {
	sqrtMap := getSqrtMap()

	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(i, "\t", sqrtMap[i])
	}
}
