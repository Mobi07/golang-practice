package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Println("Letter:", string(i))
		time.Sleep(time.Millisecond * 500)
	}
}

func main(){

	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}