package main

import (
	"fmt"
)

// can only have one func main in a Go program
// it is the entry point - the thing that will run

// BUT WHY USE POINTERS?
// YOU DONT ALWAYS WANT TO BE MAKING COPIES OF THINGS

type Kitten struct {
	Name string
}

func (k Kitten) SetName(name string) {
	k.Name = name
}

func (k Kitten) GetName() string {
	return k.Name
}

//  NOW K IS JUST A REFERENCE TO KITTY, NOT A COPY OF IT
// WITH THE ABOVE ONES, YOU GET A NEW K EVERY TIME. WITHOUT THE *, WHEN YOU CALL GETNAME YOU ARE ACTUALLY GETTING A BRAND
// NEW COPY OF KITTY.
// YOU DON'T NEED TO WORRY ABOUT THIS WITH TRUE OBJECT ORIENTED LANGUAGES
// IN GO YOU'RE ACTUALLY ATTACHING AN OBJECT TO A METHOD
// REMEMBER THE DIFFERENCE BETWEEN HAVING A STAR AND NOT HAVING A * IS THE DIFFERENCE BETWEEN HAVING A COPY AND A REFERENCE.
// WITH A REFERENCE YOU CAN CHANGE/MUTATE 'NAME'
// WITH A COPY YOURE NOT MUTATING THE SAME INSTANCE THAT YOU THINK YOU ARE.
func (k *Kitten) SetNameRef(name string) {
	k.Name = name
}

func (k *Kitten) GetNameRef() string {
	return k.Name
}

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
	kitty := Kitten{}
	kitty.SetName("Mr Tiggles")
	fmt.Println(kitty.GetName()) // THIS PRINTS NOTHING. WHY?
	kitty.SetNameRef("Mr Someone")
	kitty.SetName("who am i?")
	fmt.Println(kitty.GetNameRef()) // THIS PRINTS OUT THE NAME.

	kitty2 := &kitty
	kitty2.SetName("Mr Tom")
	fmt.Println(kitty.GetName())
}
