package main

import "fmt"

func main() {
	fmt.Println(name("Jafa"))
}

// Name prints teh user's name as
func name(s string) string {
	return fmt.Sprint("My name is ", s)
}
