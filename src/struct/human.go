package main

import (
	"fmt"
)

type Human struct {
	firstname string
	lastname  string
	age       int
}

func (h Human) eat() {
	fmt.Printf("%v eat.", h.firstname)
}

func NewFullHuman(f string, l string, a int) Human {
	return Human{f, l, a}
}

func human() {
	h1 := Human{firstname: "Nattawut"}
	h1.eat()
}
