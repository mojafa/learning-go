package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)
	jm := []string{"jenny", "money", "cash", "martini"}
	fmt.Println(jm)

	xp := [][]string{jb, jm}
	fmt.Println(xp) }
