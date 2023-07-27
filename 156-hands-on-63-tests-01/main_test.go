package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(1, 2)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}

}
