package main

import "fmt"

/*

Abstraction = hiding actual implementation and showcasing its behavior only.
In Go this is achieved using interfaces.

*/

type Payment interface {
	Pay(amount int)
}

type CreditCard struct {}

func (c CreditCard) Pay(amount int){
	fmt.Println("Paid using credit card: ", amount)
}

type Upi struct {}

func (u Upi) Pay(amount int) {
	fmt.Println("Paid using UPI: ", amount)
}

func ProcessPayment(p Payment){
	p.Pay(100)
}

func main(){
	cc := CreditCard{}
	ProcessPayment(cc)  
}