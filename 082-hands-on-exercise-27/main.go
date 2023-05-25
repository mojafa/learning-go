package main

import "fmt"

/*
create a for loop that uses break statement
● increment or decrement the variable being checked as a condition in the loop
*/
func main() {
	y := 20
	for {
		if y > 0 {
			break
		}

		// fmt.Printf("y is %v \t\t  \n", y)
		fmt.Println(y)
		y--

	}
}
