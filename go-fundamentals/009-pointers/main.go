package main

import "log"

func main() {

	myString := "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after change using pointer myString is ", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue

}
