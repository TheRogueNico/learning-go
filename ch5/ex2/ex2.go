package main

import (
	"fmt"
	"io"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer f.Close()
	data := make([]byte, 4096)
	length := 0
	for {
		count, err := f.Read(data)
		length += count
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
	}

	return length, err
}

func main() {
	length, err := fileLen("alice.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(length)
}
