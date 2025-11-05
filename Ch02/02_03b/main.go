package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
	fmt.Printf("type of str %T\n", name)
	fmt.Println("what's your age, in years?", name)
	age, _:=reader.ReadString('\n')
	fmt.Printf("type of age %T\n", age)
	ageAsNumber,  err := strconv.ParseFloat(strings.TrimSpace(age),64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your age is", ageAsNumber)
		fmt.Printf("type of ageAsNumber %T\n", ageAsNumber)
	}
}
