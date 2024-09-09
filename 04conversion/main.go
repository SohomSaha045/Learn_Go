package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza House")
	fmt.Println("Please rate our Pizza from 1 - 10")
	var reader = bufio.NewReader(os.Stdin)
	rating, err := reader.ReadString('\n')
	fmt.Println("Rating is :", rating)
		up, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
		if err != nil {
			fmt.Println("Error")
		} else {
			
			up=up+1
			fmt.Println("Rating is :", up)
		}
}
