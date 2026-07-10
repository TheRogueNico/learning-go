package main

import "fmt"

func generateNum() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	produce := func(ch chan<- int) {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}

	go produce(ch1)
	go produce(ch2)

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("goroutine 1:", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("goroutine 2:", v)
		}
	}
}

func main() {
	generateNum()
}
