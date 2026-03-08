package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(){
		fmt.Println("Sending value...")
		ch <- 5
		fmt.Println("Sent successfully")
	}()

	value := <-ch
	fmt.Println("value is:", value)
}