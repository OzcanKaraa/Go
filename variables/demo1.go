package variables

import "fmt"

func Demo1() {
	var metin string = "Hello Go"
	fmt.Println(metin)

	fmt.Println()
	fmt.Println("Hello World")

	var kdv int = 100
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25.24
	fmt.Println(kdv3)
	fmt.Printf("Veri Tipi : %T", kdv3) //Type

	var durum bool
	fmt.Println(durum)

	var isim1 string = "Özcan"
	var isim2 string = "Özcan"
	durum = isim2 == isim1
	fmt.Println(durum)
	fmt.Println(!durum)
}
