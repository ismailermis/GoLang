package main

import "fmt"

/*
Defer deyimi kendisini çevreleyen fonksiyon dönene kadar fonksiyonun çalışmasını erteler.
Ertelenmiş fonksiyon çağrısının argumanları anında değerlendirilir
fakat fonksiyon çağrısı kendisini çevreleyen fonksiyon dönmeden işleme sokulmaz.
deferAna işlevin yürütülmesinin sonunda veya ana işlev
returndeyime ulaştığında bir işlevin çalıştırılmasını sağlayan Go'daki bir anahtar sözcüktür .

*/
var isConnected = false

func main() {
	fmt.Println("================= ÖRNEK 1 ==================")
	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Printf("Connection open :%+v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection open :%+v\n", isConnected)
	fmt.Println("================== ÖRNEK 2=================")
	fmt.Println("main started")

	defer sayDone()

	fmt.Println("main finished")

}

func databaseProcessing() {
	Connect()
	fmt.Println("Defering disconnect!")
	defer disConnect()
	fmt.Printf("Connection open: %v\n", isConnected)
	fmt.Println("Do something")
}

func Connect() {
	isConnected = true
	fmt.Println("Connected database")
}
func disConnect() {
	isConnected = false
	fmt.Println("DisConnected!")
}

func sayDone() {
	fmt.Println("I am done")
}
