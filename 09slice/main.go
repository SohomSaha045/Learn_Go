package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice")
	
	var fruitList =[] string{"Apple","Mango"}
	fmt.Printf("%T \n",fruitList)
	fruitList=append(fruitList, "Tomato","Banana")
	fmt.Println(fruitList)
	fruitList=append(fruitList[1:])
	//[Mango Tomato Banana]
	//start is inclusive ==1
	fmt.Println(fruitList)
	fruitList=append(fruitList[1:3])
	//[Tomato]
	//end is not inclusive <3
	fmt.Println(fruitList)


	scores := make([]int,4)
	scores[0]=123
	scores[1]=300
	scores[2]=203
	scores[3]=23

	scores=append(scores, 409,900)
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))


	// how to remove a value from slice based on index

	var car=[] string {"Honda","Audi","Maruti","Benz"}
	fmt.Println(car)
	 var index = 2
	car=append(car[:index],car[index+1:]...)
	fmt.Println(car)


}
