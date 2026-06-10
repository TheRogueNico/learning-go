package main

import "fmt"

func prefixer(prefix string) func(string) string {
	innerPrefixer := func(inPrefix string) string {
		return prefix + " " + inPrefix
	}

	return innerPrefixer
}

func main() {
	helloPrefix := prefixer("Hello")

	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
