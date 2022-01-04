package main

import (
	"fmt"
)

func main() {

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

}
