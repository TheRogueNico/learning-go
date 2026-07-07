package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}

	greet2 := greetings[:2]
	greet3 := greetings[1:4]
	greet4 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(greet2, greet3, greet4)
}
