package main

import "fmt"

func main() {

	fmt.Println("Hello gophers")
	fmt.Println("❌")

	x := 42
	y := "James Bond"

	if x == 42 {
		fmt.Println("001")
	} else if x == 41 {
		fmt.Println("002")

	} else {
		fmt.Println("003")
	}
	if y == "James Bond" {
		fmt.Println("004")
	} else {
		fmt.Println("005")
	}

}
