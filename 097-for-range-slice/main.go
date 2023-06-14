package main

import "fmt"

func main() {
	slice := []string{"hello", "world"}
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	for i, v := range slice {
		fmt.Printf("%v - %v\n", i, v)
	}
	//i index
	// using blank identifier to not use a returned value
	for _, v := range slice {
		fmt.Printf(" %v\n", v)
	}

	//accessing valies by index position
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	fmt.Println(len(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
