package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	strx := 55
	println("float  %2f", float64(strx))
	//Console giriş için bufio paketi kullanılıyor
	//console dan geriye string ve err döner
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
	fmt.Println("Enter Number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // stringin sağındaki ve soundaki boşlukları siler. Boşluk varsa cast edemez.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
