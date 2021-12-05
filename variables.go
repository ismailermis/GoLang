package main

const aciklama = "Merhaba"
const pi = 3.14159

var (
	degisken1 = "Değişken1"
	degisken2 = "Değişken2"
)

func main() {
	/*
		var message string
		message = "Merhaba Go!"
		var message string = "Merhaba Go!"
		var message = "Merhaba GO!"
		var message = "Merhaba GO!"
		var a, b, c int
		var a, b, c int = 3, 4, 5
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		var message = "Merhaba Go!"
		var a, b, c = 3, true, 4.5
		fmt.Println(message, a, b, c)
		var k,o string="abc","xyz"
		var p=42
		var s,b="xyz",true
		// kısa değişken tanımlaması
		u:=55 // var keywordü kullanmamak için := bu opertör kullanılır.
		v,n := "abc",true
		message := "Merhaba Go!"
		a,b,c:=3,true,"test"
		a := "Metin"
		b := 'M' // burada sadece tek karakter girilebilir çünkü Char olamsını bekliyor
		c := `Metin`
		functionların dışında := şeklinde değiken tanımlanamaz.
		Global alanda tanımlanan değişkenler kullanılmaz ise derleme hatası vermez.
	*/
	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	println(myFloat32)
	println("Complex : ", myComplex)
	println(aciklama)
	println(degisken1)
	//println(pi)
}
