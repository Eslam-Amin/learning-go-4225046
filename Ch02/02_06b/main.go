package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Dates and times ---")
	goLaunchTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go Launced at %s\n", goLaunchTime)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)
	fmt.Printf("The time now is %T\n", now)

	fmt.Printf(now.Format(time.ANSIC)+ "\n")

	// addDate (year, month, day)
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.ANSIC)+ "\n")

	// format := "Mon 2006/02-01"
	format := "Mon 2006-02-01"
	fmt.Printf(tomorrow.Format(format)+ "\n")
}
