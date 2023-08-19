package main

import "fmt"

func main() {

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(add, 1, 2)
	fmt.Println(x)

	y := doMath(subtract, 1, 2)
	fmt.Println(y)
}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func doMath(f func(int, int) int, y int, x int) int {
	return f(x, y)
}
