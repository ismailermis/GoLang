package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// if
	a := 10
	b := 10

	if b > a {
		println("Büyüktür")
	} else if b == a {
		println("Eşit")
	} else {
		println("Küçüktür")
	}
	// switch
	/*
					foo := 13
					switch {
					case foo == 1:
						println("one")
					case foo == 2:
						println("two")
					case foo == 3:
						println("three")
					default:
						println("Other")
					}

				var score float64 = 0.
				println("Enter score for your last exam:")
				fmt.Scanf("%v", &score)
				switch {
				case score <= 59:
					println("Your grade is F")
				case score <= 69:
					println("Your grade is C")
				case score <= 79:
					println("Your grade is D")
				case score <= 89:
					println("Your grade is B")
				case score <= 100:
					println("Your grade is A")
				default:
					println("Please enter a score <=100")
				}
				// break and continue
				//break"
				//i := 0
				// sonsuz döngü

					for {
						if i >= 3 {
							break
						}
						println("i nin değeri : ", i)
						i++
					}

			for i := 0; i < 7; i++ {
				// if i%2 == 0 {
				// 	continue
				// }
				if i == 3 {
					continue
				} else if i == 5 {
					break
				}
				println("Çıktı :", i)
			}

		// for while c# taki while
		sum := 1
		for sum < 10 {
			sum += sum
			println(sum)
			time.Sleep(300 * time.Millisecond)
		}
	*/
	//..
	/*
	   for döngüsünün range "aralık" formu ile bir dilim "slice" veya "mao" üzerinde dolaşır.
	   Bir dilim üzerinde dolaşırken her seferinde iki değer dönülür. İlki indis, ikincisi o indisli elemanın bir kopyası.

	   Değer dizisinin indislerini veya o indise karşılık gelen değerleri _ ataması  ile atlayabilirsiniz.
	   Yanlızca indisi kullanmak istiyorsanız ", value" kısmını tamamen çıkarınız.
	*/
	//--------------------------- Range -----------------------------------------------
	/*
		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}

		ax := [3]string{"Beşiktaş", "Fenerbahçe", "Galatasaray"}
		for x := range ax {
			fmt.Println("Array item ", x, "is ", ax[x])
		}
	*/
	//------------------------- MAP ----------------------------------------------------

	baskentler := map[string]string{"Turkey": "Ankara", "Fransa": "Paris", "Italya": "Roma", "Japonya": "Tokyo"}
	for key, val := range baskentler {
		//fmt.Println("Map item: Capital of ", key, " is ", baskentler[key])
		fmt.Println("Map item: Capital of ", key, " is ", val) // ikinci yöntem

	}

}
