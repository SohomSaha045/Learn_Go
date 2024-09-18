package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Wlecome to Session on MOD")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",r)) 
}

func greeter() {
	fmt.Println("HI there mod user")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hi how are u<h1/>"))

}
