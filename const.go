package main

type Brand string

const (
	FACEBOOK      Brand = "FACEBOOK"
	MICROSOFTICON Brand = "MICROSOFTICON"
	GOOGLE        Brand = "GOOGLE"
	DIJIBI        Brand = "1453"
)

func main() {
	// go da enum yok bunun yerine constand kullanılmasını öneriyor.
	PrintBrand(GOOGLE)
	PrintBrand(FACEBOOK)
	PrintBrand(DIJIBI)
}

func PrintBrand(brand Brand) {
	println(brand)
}
