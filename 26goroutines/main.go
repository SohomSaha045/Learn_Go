package main

import (
	"fmt"
	"net/http"
	"sync"
)

// goroutines are used for parallelisum
// familiar to thread
var signal =[] string {"test"}
var wg sync.WaitGroup //usually they are pointers
var m sync.Mutex //usually pointer

func main() {
	// go greeter("hello")
	// greeter("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://briorbridgedynamics.in",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		 go getStatusCode(web)
		 wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signal)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println("String:", s)

// 	}
// }

func getStatusCode(url string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Some thing is wrong with the url.",err)
	}else{
		m.Lock()
		signal =append(signal, url)
		m.Unlock()
		
		fmt.Printf("%d status code for %s \n", res.StatusCode, url)
	}
	defer wg.Done()

}
