package main

import (
	"errors"
	"fmt"
	"log"
)

var errNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", errNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errNorgateMath
	}
	return 42, nil
}

//see use of errors.New in std libary:
//https://golang.org/src/errors/errors.go
//https://golang.org/src/bufio/bufio.go
