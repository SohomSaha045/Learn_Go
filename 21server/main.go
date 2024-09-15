package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Doing GET operation")
	// PerformGet()
	PerformPost()
}

func PerformGet() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code : ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)

	cont, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content:", string(cont))

	// var responseString strings.Builder
	// byteCount,_:=responseString.Write(cont)
	// fmt.Println("Byte Count",byteCount)
	// fmt.Println(responseString.String())

}

func PerformPost() {
	const url = "https://reqres.in/api/users"
	//payload

	body := strings.NewReader(`
	{
	"name":"Jhon Doe",
	"job":"plumber"
	}
	`)
	response, _ := http.Post(url, "application/json", body)
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response:", string(content))
	defer response.Body.Close()
}

// func PerformFormReq() {
// 	const url = "https://reqres.in/api/users"

// 	//formData

// 	data := url
// 	data.Add("")

// }
