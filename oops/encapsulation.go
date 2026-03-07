package main

import "fmt"

/*

Encapsulation = binding data + methods together and restricting direct access to fields.
In Go this is done using exported and unexported fields.
	Capitalized → Public (exported)
	Lowercase → Private (package level)

*/

type BankAccount struct {
	balance int 
}

func (b *BankAccount) Deposit(amount int){
	fmt.Println("Amount deposited successfully: ", amount)
	b.balance += amount
}

func (b *BankAccount) Balance() int {
	return b.balance
}

func main(){
	acc := BankAccount{}
	fmt.Println("Current Balance is: ", acc.balance)
	acc.Deposit(500)
	fmt.Println("Current Balance is: ", acc.balance)
}