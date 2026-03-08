package main

import "fmt"

func sendData(ch chan string){
	ch <- "Hello from sendData goroutine"
}

func main() {
	ch := make(chan string)
	go sendData(ch)

	message := <-ch
	fmt.Println("Message:", message)
	
}
