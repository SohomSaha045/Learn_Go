package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in Go")
	temp()
	// adder(3,4)
	fmt.Println("Add:",adder(5,7))
	// x:=[]int {2,3,4,5,6,7}
	// fmt.Println("Pro Add:",proAdder(5,7,9,10,8,10))
	prores,message:=proAdder(5,7,9,10,8,10)
	fmt.Println("Pro Add:",prores,"Message: ",message)

}
func adder(x int,y int) int {
	ans:=x+y
	return ans
	
}
func proAdder(val ...int)(int,string){
	sum:=0
	for _,v := range val{
		sum=sum+v
	}
	return sum,"Hello fine"
}
func temp(){
	fmt.Println("Hello from Temp Func")
}