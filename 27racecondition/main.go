package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")
	 w:= &sync.WaitGroup{}
	 mut :=&sync.Mutex{}
	var score=[]int {0}
	w.Add(3)
	go func(w *sync.WaitGroup,m *sync.Mutex){

		fmt.Println("One Routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		w.Done()
	}(w,mut)
	go func(w *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Two Routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		w.Done()

	}(w,mut)
	go func(w *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Three Routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		w.Done()

	}(w,mut)

	w.Wait()
    fmt.Println(score)
}
