package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
	//no inheritance to go lang

	sohom :=User{"Sohom","ss@go.dev",true,30}
	fmt.Println(sohom)

}
//capital meaning is Public 
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
