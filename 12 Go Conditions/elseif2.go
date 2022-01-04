package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 10

	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}

}
