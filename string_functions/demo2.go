package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isimSoyisim := "OzcanKara"
	fmt.Println(s.HasPrefix(isimSoyisim, "Ozc")) //HasPrefix:On Ek
	fmt.Println(s.HasSuffix(isimSoyisim, "ara")) //HasSuffix:Son Ek

	fmt.Println(s.Index(isimSoyisim, "Ka")) //Kacinci Index

	harfler := []string{"o", "z", "c", "a", "n", "K", "a", "r", "a"}
	fmt.Println(s.Join(harfler, "*")) //Join ile Birlestirme islemi

	harfler2 := []string{"o", "z", "c", "a", "n", "K", "a", "r", "a"}
	sonuc := (s.Join(harfler2, "-")) //Join
	fmt.Println(sonuc)

	harfler3 := []string{"o", "z", "c", "a", "n", "K", "a", "r", "a"}
	sonuc1 := (s.Join(harfler3, ""))            //Join
	fmt.Println(s.Replace(sonuc1, "", "+", -1)) //Replace:Karakter degistirme  -1 Tumunu degistir
	fmt.Println(s.Replace(sonuc1, "", "+", 3))  //En fazla 3 karakter degistir
	fmt.Println(s.Replace(sonuc1, "", "+", 1))

	//Split: Join Tam Tersi Belli formata gore ayirma parcalama yontemi
	fmt.Println(s.Split(sonuc1, "-"))

	fmt.Println(s.Repeat(sonuc1, 5)) //Repeat:Tekrarlama Kopyasini yazdirma

}
