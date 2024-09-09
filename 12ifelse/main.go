package main

import "fmt"

func main() {
	fmt.Println("If else")
	count:=19

	var result string

	if count<10{
		result="Not regular User"
	} else if count<20{
		result="Regular User"
	} else{
		result="Mad User"
	}
	fmt.Println(result)

	if 9%2==0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	if num:=9;num%2==0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
}
