package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps")

	pl := make(map[string]string)

	pl["C++"] = "fast"
	pl["Java"] = "medium"
	pl["Go"] = "very fast"

	fmt.Println(pl)
	fmt.Println(pl["Java"])

	delete(pl, "Java")
	fmt.Println(pl)


	//loop through map

	for key,value := range pl{
		fmt.Println(key,":",value)
	}

}
