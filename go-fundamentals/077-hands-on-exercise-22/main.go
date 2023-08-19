package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	if x < 4 && y < 4 {
		fmt.Printf("x=%v and y=%v are both less than 4\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("x=%v and y=%v are both greater than 6\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x=%v is greater than or equal to 4 and less than or equal to 6\n", x)
	} else if y != 5 {
		fmt.Printf("y=%v is not equal to 5\n", y)
	} else {
		fmt.Printf("none of the previous cases were met")

	}
}

/*
Create 2 random ints between 0 inclusive and 10 exclusive
○ assign them to variables with the identifiers x and y ● Print their values
if statement to print the correct description xandyarebothlessthan4
x and y are both greater than 6
x is greater than or equal to 4 and less than or equal to 6 yisnot5
*/
