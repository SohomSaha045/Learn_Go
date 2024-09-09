package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"Sunday", "Tuesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for d:=range days{
	// 	fmt.Println(days[d])
	// }

	// for index,day := range days{
	// 	fmt.Println("Index-",index,"Day-",day)
	// }

	val := 1

	for val < 10 {

		if val == 4 {
			val++
			continue
		} else if val == 9 {
			break
		}else if val==7{
			goto Help
		}

		fmt.Println("Val:", val)
		val++
	}
Help:
	fmt.Println("Help")
}
