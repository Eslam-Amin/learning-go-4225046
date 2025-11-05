package main

import (
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something")
	value1 := 5
	value2 := 10
	value3 := 56
	values := make([]int, 0, 5)
	values = []int{12,15,16,22,90}
	sum, count, average := addAllValues(values...)
	fmt.Printf("The sum is %v.\n", sum)
	fmt.Printf("The count is %v.\n", count)
	fmt.Printf("The average is %v.\n", average)
	
	sum, count, average = addAllValues(value1, value2, value3, 125)
	fmt.Printf("The sum is %v.\n", sum)
	fmt.Printf("The count is %v.\n", count)
	fmt.Printf("The average is %v.\n", average)

}

func addValues(value1, value2 int) int {
	return value1 + value2
}

	func addAllValues(values ...int) (int, int, float64) {
		sum := 0
		for _, v := range values {
			sum += v
		}
		count := len(values)
		average := float64(sum) / float64(count)
		return sum, count, average
	}
