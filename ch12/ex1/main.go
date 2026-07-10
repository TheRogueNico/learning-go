package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	numWrite := func() {
		defer wg.Done()
		for i := range 10 {
			nums <- i
		}
	}

	go numWrite()
	go numWrite()

	go func() {
		wg.Wait()
		close(nums)
	}()

	var consumer sync.WaitGroup
	consumer.Add(1)

	go func() {
		defer consumer.Done()
		for n := range nums {
			fmt.Println(n)
		}
	}()

	consumer.Wait()
}
