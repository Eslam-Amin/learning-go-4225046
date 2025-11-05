package main

import (
	"fmt"
)

func main() {
	// var colors [4]string
	// colors[0] = "Red"
	// colors[1] = "Green"
	// colors[2] = "Blue"
	var colors [4]string = [4]string{"Red", "Green", "Blue"}
	colors[3] = "Black"
	fmt.Println("Colors:", colors)

	var number = [5]int{5, 1, 2, 4}
	fmt.Println("Numbers:", number)
	fmt.Println("len(Numbers):", len(number))
	fmt.Println("len(Colors):", len(colors))
}
