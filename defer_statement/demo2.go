package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	if sayi%2 == 0 {
		return "Çift Sayidir."
	}

	if sayi%2 != 0 {
		return "Tek Sayidir."
	}
	return "Belli degil"
}

func Test() {
	sonuc := Demo2(11)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("İslem Bitti")
}
