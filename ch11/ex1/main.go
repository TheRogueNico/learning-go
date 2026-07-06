package main

import (
	"embed"
	"fmt"
	"os"
)

//go:embed rights
var humanRights embed.FS

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<language>")
		os.Exit(1)
	}

	text, err := humanRights.ReadFile("rights/" + os.Args[1] + "_rights.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	println(string(text))
}
