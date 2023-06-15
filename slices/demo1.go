package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "Ozcan"
	isimler[1] = "Derin"
	isimler[2] = "Salih"

	fmt.Println(isimler)
}
