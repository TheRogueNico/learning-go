package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	total, i := 0, 0
	reason := ""

loop:
	for {
		select {
		case <-ctx.Done():
			reason = "Timeout"
			break loop
		default:
			n := rand.IntN(100_000_000)
			total += n
			i++
			if n == 1234 {
				reason = "Reached 1234"
				break loop
			}
		}
	}

	fmt.Println("Reason: ", reason, "\nIteration:", i, "\nTotal: ", total)
}
