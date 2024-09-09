package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	x:=45
	var one *int=&x

	fmt.Println(*one)
}
