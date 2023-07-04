package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, "and i'm ", p.age, "years old")
}
func main() {
	p1 := person{
		first: "James",
		age:   30,
	}
	p2 := person{
		first: "Jenny",
		age:   30,
	}
	p1.speak()
	p2.speak()

}

/*
Create a user defined struct with
the identifier “person”
the fields:
first
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
