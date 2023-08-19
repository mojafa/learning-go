package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	println(i)
	// }

	for i := 0; i < 100; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("iteration %v \t x=%v and y=%v are both less than 4\n", i, x, y)

		case x > 6 && y > 6:
			fmt.Printf("iteration %v \t x=%v and y=%v are both greater than 6\n", i, x, y)
		case x >= 4 && x <= 6:
			fmt.Printf("iteration %v \t x=%v is greater than or equal to 4 and less than or equal to 6\n", i, x)
		case y != 5:
			fmt.Printf("iteration %v \t y=%v is not equal to 5\n", i, y)
		default:
			fmt.Printf("iteration %v \t none of the previous cases were met", i)
		}

	}
}

/*
 there are two parts ot this hands on exercise
○ create a program that has a loop that prints every number from 0 to 99
○ modify the program from the previous hands on exercise to run 100 times
*/
