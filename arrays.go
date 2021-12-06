package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 5} // Eğer dizi tam doldurulmazsa kalan kısımlara 0 basılır.Number of numbers : [1 2 3 5 0]
	fmt.Println(numbers)
	fmt.Println("Number of numbers :", numbers, len(numbers))
	myArray := [...]int{1, 2, 3, 5, 56, 369, 4865} // otomatik boyutlandırma yapar [...]
	myArray[3] = 222
	fmt.Println("Number of myArray :", myArray, len(myArray), " Kapasitesi :", cap(myArray))

	for i, value := range numbers {
		fmt.Println("Number of ", i, " Degeri : ", value)
	}

}
