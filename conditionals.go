package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 100 {
		fmt.Println("X is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("X is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	n := 3

	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	switch {
	case n > 100:
		fmt.Println("n is very big")
	case n > 10:
		fmt.Println("n is  big")
	default:
		fmt.Println("n is  small")
	}
}
