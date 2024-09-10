package main

import (
	"fmt"
	"net/url"
)

const URL="https://lco.dev:8000/learn?course=reactjs&paymentid=hsdsdsjg"


func main() {
	fmt.Println("Welcome to URL Section in GO")
	// fmt.Println(URL)
	
	
	//parsing url

	res,err:=url.Parse(URL)

	if err!=nil{
		panic(err)
	}
	fmt.Println("Result:",res)
	fmt.Println("Result Scheme:",res.Host)
	fmt.Println("Result Query:",res.RawQuery)
	fmt.Println("Result PORT:",res.Port())
	fmt.Println("Result Path:",res.Path)

	params:= res.Query()
	fmt.Printf("The Type of Params are %T \n",params)
	// fmt.Println(params["course"])
	// fmt.Println(params["paymentid"])

	for _,val := range params{
		fmt.Println(val)
	}



	parts:= &url.URL{
		Scheme: "https",
		Host: "makdis.dex",
		Path: "/top",
		RawQuery: "user=90888sh",
	}

	second:=parts.String()
	fmt.Println(second)
}
