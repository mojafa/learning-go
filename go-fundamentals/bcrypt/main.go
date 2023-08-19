package main

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password 123`
	cleanedPass := strings.ReplaceAll(s, " ", "")
	bs, err := bcrypt.GenerateFromPassword([]byte(cleanedPass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPass := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You are logged in")
}
