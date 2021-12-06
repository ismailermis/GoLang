package main

// Ananonim fonksiyonlar tek bir kez çalıştırılırlar.
// çalışıp işi bittikten sonra hafızadan silinirler.

func main() {
	addfunc := func(terms ...int) (numterms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numterms = len(terms)
		return
	}
	numterms, sum := addfunc(2, 5)
	println("Added :", numterms, " terms for a total of ", sum)
}
