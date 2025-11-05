package main

import (
	"fmt"
)

func main() {
	fname,lname:="",""
	// str1 := "The quick red fox"
	// fmt.Println("typeof str1 ", reflect.TypeOf(str1))
	// str2 := "jumped over"
	// str3 := "the lazy brown dog."
	fmt.Println("What's your name?")
	fmt.Scanln(&fname, &lname)
	fmt.Println("Hello from Go!,",fname,lname)
	// fmt.Println(str1)
	// input:=fmt.Scanf()
	stringLength , err := fmt.Println(fmt.Scanf(""))
	if err == nil{
		fmt.Println("String length:", stringLength)
	}else{
		fmt.Println(err)
	}
	fmt.Printf("Data type of err : %T\n", err)
	fmt.Printf("Data type of fname : %T\n", fname)
}
