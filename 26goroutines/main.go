package main

import (
	"fmt"
	"time"
)

//goroutines are used for parallelisum
//familiar to thread
func main() {
	go greeter("hello")
	greeter("world")
}
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3*time.Millisecond)
		fmt.Println("String:", s)

	}
}
