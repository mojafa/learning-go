package main

import "fmt"

func main() {
	y()
	x := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	x()
}

var y = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
