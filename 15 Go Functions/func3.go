package main

import (
	"fmt"
)

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
	familyName("liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)
}
