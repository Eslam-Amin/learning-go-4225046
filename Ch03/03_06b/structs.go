package main

import "fmt"

func structs() {

	// This is how your code will be called.
	// Your answer should be a slice of Color objects.
	// You can edit this code to try different testing cases.
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)

	fmt.Println(colors)
}

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	colors := make([]Color,0, len(colorNames))
	for i := range colorNames {
		colors = append(colors, Color{Name: colorNames[i], Hex: hexValues[i]})
	}
	return colors
}

type Color struct {
	Name string
	Hex  int
}