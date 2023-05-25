package main

import "fmt"

func main() {
	//create a for loop using only a condition
	y := 20
	// for {
	// 	fmt.Printf("y is %v \t\t  \n", y)
	// 	if y < 2 {
	// 		break
	// 	}
	// 	y--
	// }

	for y > 0 {
		fmt.Printf("y is %v \t\t  \n", y)
		y--
	}
}

/*
create a for loop using only a condition
● increment or decrement the variable being checked in the condition
*/
