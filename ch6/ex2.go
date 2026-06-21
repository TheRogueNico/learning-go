package main

import "fmt"

func UpdateSlice(slice []string, new string) {
	slice[len(slice)-1] = new

	fmt.Println(slice)
}

func GrowSlice(slice []string, new string) {
	slice = append(slice, new)

	fmt.Println(slice)
}

func main() {
	s := []string{"Ant", "Banana", "Coconut"}
	n := "Dragon"

	fmt.Println(s)
	UpdateSlice(s, n)
	fmt.Println(s, "\n-----")

	fmt.Println(s)
	GrowSlice(s, n)
	fmt.Println(s)
}
