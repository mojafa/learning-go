package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	fmt.Println("-------------")
	//variadic parameter
	xi = append(xi, 5, 6, 7, 8, 9, 10)

	fmt.Println(xi)
	fmt.Println("-------------")
}
