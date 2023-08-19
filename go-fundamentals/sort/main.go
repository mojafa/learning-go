package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 5, 7, 8, 42}
	xs := []string{"James", "Doe", "Jenny", "Jone"}

	fmt.Println(xi)
	sort.Ints(xi)

	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

}
