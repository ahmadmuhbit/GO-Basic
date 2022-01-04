package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func main() {
	familyName("liam")
	familyName("Jenny")
	familyName("Anja")
}
