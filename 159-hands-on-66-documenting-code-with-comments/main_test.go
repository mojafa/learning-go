package main

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	got := name("Jafa")
	want := "My name is Jafa"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}

}
