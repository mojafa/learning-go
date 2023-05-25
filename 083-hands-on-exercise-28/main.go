package main

import "fmt"

/*
use modulus and the continue statement in a loop to print all ODD numbers
● joke about the programmer and infinite loops
*/
func main() {
	for i := 0; i < 14; i++ {
		// if i%2 == 1 {
		// 	fmt.Println(i)
		// 	continue
		// }
		if i%2 != 0 {
			fmt.Println(i)
			continue
		}
	}
}
