package main

import (
	"fmt"
)

func main() {

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // Menentukan urutannya
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // Loop tanpa urutan
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // Loop dengan urutan yang ditentukan
		fmt.Printf("%v : %v, ", element, a[element])
	}

}
