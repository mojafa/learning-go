package main

import "fmt"

func main() {
	fmt.Println(name("Jafa"))

}

func name(s string) string {
	return fmt.Sprint("My name is ", s)
}
