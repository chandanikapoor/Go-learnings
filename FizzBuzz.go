package main

import (
	"fmt"
)

/* print numbers between 1 to 20
if num is divisible by 3 , print fizz,
if num is divisible by 5, print buzz,
if num is divisible by 3 & 5, print fizz buzz
otherwise print the number
*/
func main() {
	for i := 0; i < 21; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
