package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer")
	defer fmt.Println("Hello World1")
	fmt.Println("Hello World no defer")
	defer fmt.Println("Hello World2")
	temp()
}
func temp()  {
	for i:=0;i<6;i++{
		defer fmt.Println(i)
	}
	
}
