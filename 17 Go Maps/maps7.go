package main

import (
	"fmt"
)

func main() {

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a["brand"] // Memeriksa kunci yang ada dan nilainya
	val2, ok2 := a["color"] // Memeriksa kunci yang tidak ada dan nilainya
	val3, ok3 := a["day"]   // Memeriksa kunci yang ada dan nilainya
	_, ok4 := a["model"]    // Hanya memeriksa kunci yang ada dan bukan nilainya

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

}
