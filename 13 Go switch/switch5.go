package main

import (
	"fmt"
)

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a is one")
	case "b":
		fmt.Println("a is b")
	}
}
