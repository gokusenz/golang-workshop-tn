package main

import (
	"fmt"
)

func main() {
	x := make(map[string]string)
	x["firstname"] = "Nattawut"

	fmt.Printf("%T => %v\n", x, x)

	if data, ok := x["firstname"]; ok {
		fmt.Println(data, ok)
	}
}
