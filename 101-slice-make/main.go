package main

import "fmt"

func main() {
	// making via composite literals
	si := []string{"a", "b", "c", "d", "e"}
	fmt.Println(si)

	//making via make
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(xi)
	fmt.Println("------------")
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
