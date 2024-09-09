package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Ludo Using switch Case")

	rand.Seed(time.Now().UnixNano())
	dice:=rand.Intn(6) + 1
	fmt.Println("Value of Dice",dice)

	switch dice {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fallthrough
	default:
		fmt.Println("Other than One Two")
	
		
	}
}
