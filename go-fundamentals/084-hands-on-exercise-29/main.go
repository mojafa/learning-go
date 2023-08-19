package main

import "fmt"

/*
create a loop that runs 5 times
● nest a loop within the first loop that also prints 5 times
● print something in each loop to illustrate what is occuring
*/
func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}
}
