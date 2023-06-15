package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "Ozcan"
	isimler[1] = "Derin"
	isimler[2] = "Salih"
	isimler = append(isimler, "Berfin") // Yeni Eleman ekleme

	fmt.Println(isimler)
	fmt.Println(len(isimler))
}