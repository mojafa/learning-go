package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 0
	y := 100
	// fmt.Println(rand.Intn(y-x)+x, "is a random number between", x, "and", y)
	z := rand.Intn((y - x) + x)
	fmt.Printf("z=%v \n", z)
	//conditionals
	if z >= 0 && z <= 100 {
		fmt.Printf("z is %v and it's less than 100\n", z)
	} else if z >= 101 && z <= 200 {
		fmt.Printf("z is %v and that is between 101 and 200\n", z)
	} else {
		fmt.Printf("z is %v and that is between 201 and 250\n", z)
	}
	a := rand.Intn(350)
	if a <= 100 {
		fmt.Printf("a is %v and that is between 0 and 100\n", a)
	} else if a >= 101 && a <= 200 {
		fmt.Printf("a is %v and that is between 101 and 200\n", a)
	} else if a >= 201 && a <= 250 {
		fmt.Printf("a is %v and that is between 201 and 250\n", a)
	} else {
		fmt.Println("a is greater than 250")
	}

	/* 	rand.Seed(time.Now().UnixNano())
	   	min := 0
	   	max := 30
	   	fmt.Println(rand.Intn(
	   		(max - min + 1) + min,
	   	)) */

	//test inclusinve vs  exclusive Intn [0,n) 0 to n-1]
	// 0 is inlcuded n is not included!

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

}

/* create a program that uses both SEQUENTIAL and CONDITIONAL control flow.
Your program should do the following:
○ create a random int between 0 and 250
○ store the value of that int in a variable with the identifier of x
■ func Intn(n int) int
	● rand.Intn()
○ print out the name and value of the variable
○ use an if statement to do the following
■ if the ● value is between 0 and 100, print between 0 and 100
■ if the ● value is between 101 and 200, print between 101 and 200
■ if the ● value is between 201 and 250, print between 201 and 250
● re: inclusive, exclusive – does rand.Intn()
○ include zero in its output?
○ include 250 in its output?
○ show this in code using the numbers 0 - 3

*/
