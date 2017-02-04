package main

import (
	"fmt"

	"github.com/katsuraku/wwg/animals"
)

// can only have one func main in a Go program
// it is the entry point - the thing that will run

func main() {
	fmt.Println("yo")
	a := 5
	b := a
	// b doesn't really equal a. b is a copy of a - at that point in time.
	// it gets the value copied, and it's kind of flattened. no history. it HAS that value.
	// it doesn't remember that it is a copy. so if you subsequently change the a, it has no effect on b.
	a = 7

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	// POINTERS

	c := 5
	d := &a
	c = 7

	fmt.Println("c is ", c)
	fmt.Println("d's memory address/pointer location/reference is ", d) // that prints out the reference, the memory address
	fmt.Println("d, dereferenced using * is ", *d)                      // dereferenced - looks through the pointer and gives the value when you look wherever the address is pointing you to.

	// BUT WHY USE POINTERS? CHECK THIS OUT
	kitty := animals.Kitten{}
	kitty.SetName("Mr Tiggles")
	fmt.Println(kitty.GetName()) // THIS PRINTS NOTHING. WHY?
	kitty.SetNameRef("Mr Someone")
	kitty.SetName("who am i?")
	fmt.Println(kitty.GetNameRef()) // THIS PRINTS OUT THE NAME.

	kitty2 := &kitty
	kitty2.SetName("Mr Tom")
	fmt.Println(kitty.GetName())
}
