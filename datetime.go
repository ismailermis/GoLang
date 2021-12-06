package main

import (
	"fmt"
	"strings"
	"time"
)

/*unix datetime  unix sistemlerde sistem tarihinin formatina verlen isim.
 "unix time" 1/1/1970 00:00 tarihinden itibaren gecen saniye sayisina denk du$en bir integer'dir..
Unix, 1 Ocak 1970'den bu yana geçen süreyi saniye sayısı olarak saklar. Bu, Linux'un da yaptığı anlamına gelir.

*/
func main() {
	r := strings.NewReplacer(
		"January", "Ocak",
		"February", "Şubat",
		"March", "Mart",
		"April", "Nisan",
		"May", "Mayıs",
		"June", "Haziran",
		"July", "Temmuz",
		"August", "Ağustos",
		"September", "Eylül",
		"October", "Ekim",
		"November", "Kasım",
		"December", "Aralık")
	/*
		fmt.Printf("Şuan ki zaman :%v\n", time.Now().Unix())
		now := time.Now()

		fmt.Println("Year:", now.Year())
		fmt.Println("Month:", now.Month())
		fmt.Println("Day:", now.Day())
		fmt.Println("Hour:", now.Hour())
		fmt.Println("Minute:", now.Minute())
		fmt.Println("Second:", now.Second())
		fmt.Println("Nanosecond:", now.Nanosecond())
		fmt.Println("===================================")
		// Defining t for UTC method
		t := time.Now()

		// Calling UTC method
		utc := t.UTC()

		// Prints output
		fmt.Printf("%v\n", utc.Local())

		time := utc.Format("2 January 2006")
		outputString := r.Replace(time)
		fmt.Println(outputString)
	*/
	fmt.Println("===================================")
	t1 := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("%s\n", t1)
	fmt.Println("===================================")
	fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
	fmt.Println("===================================")
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"))
	fmt.Println("===================================")
	time1 := time.Now().Format("2 January 2006 15:04:05")
	outputString := r.Replace(time1)
	fmt.Println(outputString)
	fmt.Println("===================================")
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().AddDate(0, 0, 1))
	fmt.Println("=============> DIFF <==============")

	x := fmt.Println
	diff := time.Now().Sub(t1)
	x(diff)
	x(diff.Hours())
	x(diff.Minutes())
	x(diff.Seconds())

}
