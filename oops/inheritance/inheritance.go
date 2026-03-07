package main

import "fmt"

/*

Inheritance - one class to acquire the fields and methods of another class using the extends keyword.
It represents an “is-a” relationship between classes.
Go does NOT support inheritance
Instead Go uses composition (embedding).
This is often called "composition over inheritance". It represents an “has-a” relationship between structs.

*/

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Animal Speaks")
}

type Dog struct {
	Animal // Embedding the Animal struct
}

func main() {
	dog := Dog{Animal{Name: "Tommy"}}

	dog.Animal.Speak()

	// Speak method is promoted to Dog struct, so we can call it directly on dog
	dog.Speak()
}
