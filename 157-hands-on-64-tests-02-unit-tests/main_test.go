// Test files live in the same package as the code they test.
// They are named with the following convention:
// `filename_test.go` where filename is the name
// of the file that contains the code you want to test.

package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(1, 2)
	want := 3
	if got != want {
		log.Fatalf("error - want %v, got %v", want, got)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 5)
	want := 2
	if got != want {
		log.Fatalf("error - want %v, got %v", want, got)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(add, 7, 5)
	want := doMath(add, 7, 5)
	if got != want {
		log.Fatalf("error - want %v, got %v", want, got)
	}
}
