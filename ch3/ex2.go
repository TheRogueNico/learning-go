package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	char := []rune(message)

	fmt.Println(string(char[3]))
}
