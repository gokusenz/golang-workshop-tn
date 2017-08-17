package main

import (
	"fmt"
)

func main() {
	message := "Hello world"
	fmt.Printf("%s\n", message)

	i := 1
	f := 1.0
	b := true
	s := "String"

	fmt.Printf("%T => %v\n", i, i)
	fmt.Printf("%T => %v\n", f, f)
	fmt.Printf("%T => %v\n", b, b)
	fmt.Printf("%T => %v\n", s, s)
}
