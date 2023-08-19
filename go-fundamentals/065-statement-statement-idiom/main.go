package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("first")
	fmt.Println("second")

	x := 40
	y := 5

	fmt.Printf("x=%v \n y=%v \n", x, y)

	//conditionals

	if z := 2 * rand.Intn(x); z > x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)

	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}

}
