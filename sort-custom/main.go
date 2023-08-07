package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Jenny", 27}
	p3 := person{"Jone", 44}
	p4 := person{"Doe", 12}
	p5 := person{"Amar", 10}
	p6 := person{"Dann", 2}

	people := []person{p1, p2, p3, p4, p5, p6}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	//sort by name
	sort.Sort(ByName(people))
	fmt.Println(people)

}
