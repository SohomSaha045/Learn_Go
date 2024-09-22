package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SohomSaha045/mongoAPI/router"
)

func main() {
	fmt.Println("Hi from Mongo Connection")
	r:=router.Router()
	//sohomsahacse2020
	//Fv0RU1NbAcPnfe0r
	//mongodb+srv://sohomsahacse2020:<db_password>@cluster0.6dx48.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
	fmt.Println("Server is getting Started...")
	log.Fatal(http.ListenAndServe(":8080",r))
	fmt.Println("Listening to PORT 8080...")
}
