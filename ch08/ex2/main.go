package main

import "fmt"

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func Print[T Printable](t T) {
	fmt.Println(t)
}

type (
	MyInt   int
	MyFloat float64
)

func (n MyInt) String() string {
	return fmt.Sprintf("MyInt is:\t%d", int(n))
}

func (f MyFloat) String() string {
	return fmt.Sprintf("MyFloat is:\t%f", float64(f))
}

func main() {
	Print(MyInt(12))
	Print(MyFloat(0.86))
}
