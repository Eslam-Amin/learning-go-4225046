package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	}else{
		fmt.Println("the value of p is",*p)
	}
	
	
	fValue1 := 42.13
	pointerOfValue1 := &fValue1
	fmt.Println("the value of pointerOfValue1 is",*pointerOfValue1)
	
	*pointerOfValue1 = *pointerOfValue1 / 31
	fmt.Println("the value of pointerOfValue1 is",*pointerOfValue1)
	fmt.Println("the value of fValue1 is",fValue1)
	
}

