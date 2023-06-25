package pointers

import "fmt"

//Pointer :

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı :", *sayi)
}
