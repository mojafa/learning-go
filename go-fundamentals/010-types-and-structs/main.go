package main

import (
	"log"
	"time"
)

var s = "seven"

// var fname string
// var lname string
// var age int
// var dob time.Time
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "555-555-5555",
		Age:         25,
		BirthDate:   time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)

	var s2 = "six"
	log.Print(s)
	log.Print(s2)
	saySomething("Hello")
}

func saySomething(s3 string) (string, string) {
	log.Print("s from say something is: ", s)
	return s3, "World"
}
