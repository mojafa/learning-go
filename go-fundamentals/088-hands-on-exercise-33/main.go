package main

import (
	"fmt"
	"math/rand"
)

func main() {

	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t total count %v \t  x is  %v\n", i, c, x)
			c++
		}
	}

}

/*
use the "statement statement" idiom to
○ initialize x with and random int between 0 inclusive and 5 exclusive ○ ifxis3
■ print "x is 3"
● run that code 100 times
● what's the benefit of using the "statement statement" idiom?

*/
