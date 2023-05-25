package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 43,
	}
	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}

}
