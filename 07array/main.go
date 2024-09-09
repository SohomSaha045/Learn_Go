package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")
	var list [4] string
	list[0]="hi"
	list[2]="How"
	list[3]="are"
	fmt.Println("List is",list)

	var car=[2]string {"audi","honda"}
	fmt.Println(car,len(car))
}
