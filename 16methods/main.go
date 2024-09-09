package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
	//no inheritance to go lang

	sohom :=User{"Sohom","ss@go.dev",true,30,9}
	fmt.Println(sohom.Email)
	// sohom.GetStatus()
	sohom.SetEmail()
	fmt.Println(sohom.Email)


}
//capital meaning is Public 
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int //private
}

func (u User) GetStatus() {
	fmt.Println("User Status is:" , u.Status)
	fmt.Println("One Age is:" , u.oneAge)
}
func (u User) SetEmail() {
	fmt.Println("User Email is:" , u.Email)
	u.Email="test@gmail.com"
	fmt.Println("New Email is:" , u.Email)
}
