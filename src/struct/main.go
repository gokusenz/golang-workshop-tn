package main

import (
	"fmt"
)

type Human struct {
	firstname string
}

func (h Human) eat() {
	fmt.Printf("%v eat.", h.firstname)
}

func main() {
	h1 := Human{firstname: "Nattawut"}
	h1.eat()
}
