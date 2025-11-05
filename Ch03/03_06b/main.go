package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Structs ---")
	// poodle := Dog{"Poodle", 34, "Woof"} // to create a struct (similar to a class)
	// fmt.Println(poodle)
	// fmt.Printf("%+v\n",poodle) // %+v will print the struct in a more human readable format
	// poodle.Speak()
	// poodle.Sound = "Arf"
	// poodle.Speak()
	structs()
}

type Dog struct{
	Breed string
	Wight int
	Sound string

}

	func (d Dog) Speak() {
		fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
	}
