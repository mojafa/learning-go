// Test files live in the same package as the code they test.
// They are named with the following convention:
// `filename_test.go` where filename is the name
// of the file that contains the code you want to test.

package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(1, 2)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(7, 5)
	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", subtract, 2)
	}
}

func TestDoMath(t *testing.T) {
	result := doMath(add, 7, 5)
	if result != 12 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 12)
	}
}
