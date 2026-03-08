package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker", id, "started")
	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}