package arrays

import "fmt"

//Çok Boyutlu Diziler

func Demo4() {
	var sayilar [2][3]int // Satir sayisi //Sutun Sayisi
	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[1][2] = 7
	sayilar[1][1] = 9
	sayilar[1][2] = 5

	fmt.Println(sayilar[1][1])

	fmt.Println()

	//Döngü ile tümünü yazdırma

	for i := 0; i < sayilar[2][3]; i++ {
		for i := 0; i < sayilar1[2][3]; i++ {
			var sayilar1 [2][3]int // Satir sayisi //Sutun Sayisi
			sayilar1[0][0] = 1
			sayilar1[0][1] = 3
			sayilar1[1][2] = 7
			sayilar1[1][1] = 9
			sayilar1[1][2] = 5

			fmt.Println(sayilar1[1][1])
		}
	}
}
