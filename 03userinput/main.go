package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading:")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
