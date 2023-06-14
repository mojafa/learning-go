package main

import "fmt"

func main() {
	slice := []string{"hello", "world"}
	fmt.Println(slice)
}

// slices is a data dtructure with three elements:s:
// -length
// -cap
// - pointer to an underlying array

// type slice struct {
// 	array unsafe.Pointer
// 	len int
// 	cap int
// }
