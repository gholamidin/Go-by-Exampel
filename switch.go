package main

import (
	. "fmt"
	"time"
)

func main() {

	i := 2
	Println("write", i, "as")
	switch i {
	case 1:
		Println("one")
	case 2:
		Println("two")
	case 3:
		Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		Println("it is weekend")
	default:
		Println("it is a weeksday")
	}

}
