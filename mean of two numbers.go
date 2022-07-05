package main

import (
	"fmt"
)

func main() {
	/*var x int
	var y int

	x = 10
	y = 20
	*/
	/*var x float64
	var y float64

	x = 15.5
	y = 20.0

	x := 15.5
	y := 20.0
	*/

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result:%v, type of %T\n", mean, mean)
}
