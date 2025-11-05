package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	// var newInt int64 = 15
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3:= 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)
	intFloatSum := f1 + float64(i1)
	fmt.Println("Result:", intFloatSum)
	
	floatIntSum := int(f1) + i1
	fmt.Println("Result:", floatIntSum)
	product := float64(i1) * f2
	fmt.Println("Product:", product)

	
}
