package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x=%v and y=%v are both less than 4\n", x, y)

	case x > 6 && y > 6:
		fmt.Printf("x=%v and y=%v are both greater than 6\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("x=%v is greater than or equal to 4 and less than or equal to 6\n", x)
	case y != 5:
		fmt.Printf("y=%v is not equal to 5\n", y)
	default:
		fmt.Printf("none of the previous cases were met")
	}
}

/*
● Modify the previous program to use a switch statement
*/
