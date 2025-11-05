package main

import (
	"fmt"
	"sort"
)


func main() {
	// This is a slice
	// var colors = []string{"Red", "Green", "Blue"}
	// var colors = make([]string, the initial items in the slice , slice's capacity)
	
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	// colors = append(colors, "Black", "Orange", "Yellow")
	// fmt.Println(colors)
	fmt.Println("Here is the colors array", colors)
	loop:
	for {

		fmt.Println("do you want to append new color or remove?")
		fmt.Println("add => 1 \n remove => 2 \n exit => 3")
		
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var newColor string
			fmt.Println("Enter the new Color")
			fmt.Scanln(&newColor)
			colors = append(colors, newColor)
			fmt.Println("Here is the updated colors array", colors)
			case 2:
				fmt.Println("please enter the index of the color you want to remove")
				var index int
				fmt.Scanln(&index)
				colors = remove(colors, index)
				fmt.Println("Here is the updated colors array", colors)
			case 3:
					fmt.Println("Goodbye")
					break	loop
			default:
					fmt.Println("Invalid choice")
				
			}
	}
	sort.Strings(colors)
	fmt.Println("sorted Colors: ",colors)
}


func remove (slice []string, i int) []string{
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}
