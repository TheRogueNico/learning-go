package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	} // the total in the for loop dies here

	// this is the 0 total we defined at top of main
	fmt.Println(total)
}
