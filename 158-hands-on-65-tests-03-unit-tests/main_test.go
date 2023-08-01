package main

import "testing"

func TestName(t *testing.T) {
	got := name("Jafa")
	want := "My name is Jafa"
	if got != want {
		t.Errorf("error - want %v, got %v", want, got)
	}

}
