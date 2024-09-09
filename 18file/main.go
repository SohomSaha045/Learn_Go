package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handing in Go")
	content := "Hi hello how are you."
	file, err := os.Create("./myNotebook.txt")

	// if err!=nil{
	// 	fmt.Println("Error")
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is :", length)
	readfile("myNotebook.txt")
	defer file.Close()
}

func readfile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text DataBytes in File is :", databyte)
	fmt.Println("Text Data in File is :", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
