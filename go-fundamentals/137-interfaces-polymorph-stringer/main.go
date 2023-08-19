package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("the title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("this is the number ", strconv.Itoa(int(c)))
}
func main() {
	b := book{
		title: "West With the Night",
	}
	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }