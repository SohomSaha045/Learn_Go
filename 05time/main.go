package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")
	present := time.Now()
	fmt.Println(present)
	fmt.Println(present.Format(
		"01/02/2006 15:04:05 Monday"))

	create := time.Date(2020, time.December, 12, 4, 54, 23, 0, time.Local)
	fmt.Println(create)
}
