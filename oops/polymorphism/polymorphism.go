package main

import (
	"fmt"
	"math"
)

/*

Polymorphism allows the same method or object to behave differently based on the context, specially on the project's actual runtime class.
Polymorphism = same interface, different behavior.
Go achieves polymorphism through interfaces

*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	if math.Mod(c.Radius, 7) == 0 {
		return (22 * c.Radius * c.Radius) / 7
	}
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	PrintArea(rect)
	PrintArea(circle)
}
