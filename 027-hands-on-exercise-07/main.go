package main

import "fmt"

func main() {
	var a int = 747
	fmt.Printf("747 as decimal, %d \n", a)
	fmt.Printf("747 as binary, %b \n", a)
	fmt.Printf("747 as octal, %o \n", a)
	fmt.Printf("747 as hexadecimal, %x \n", a)
	fmt.Println()

	var b int = 911
	fmt.Printf("911 as decimal, %d \n", b)
	fmt.Printf("911 as binary, %b \n", b)
	fmt.Printf("911 as octal, %o \n", b)
	fmt.Printf("911 as hexadecimal, %x \n", b)

	fmt.Println()
	var c int = 90210
	fmt.Printf("90210 as decimal, %d \n", c)
	fmt.Printf("90210 as binary, %b \n", c)
	fmt.Printf("90210 as octal, %o \n", c)
	fmt.Printf("90210 as hexadecimal, %x \n", c)

}

// write a program that uses print verbs to show the following numbers
// ● 747
// ● 911
// ● 90210
// as decimal, binary, hexadecimal
