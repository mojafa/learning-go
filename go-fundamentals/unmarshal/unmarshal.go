package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	ID       int    `json:"id"`
	Customer string `json:"customer"`
	Items    []item `json:"items"`
}

type item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	s := `{"id":1,"customer":"John Doe","items":[{"id":1,"name":"iPhone 12","price":1000},{"id":2,"name":"iPhone 12 Pro","price":1200}], 
	"id":2,"customer":"Jane Doe","items":[{"id":1,"name":"iPhone 12","price":1000},{"id":2,"name":"iPhone 12 Pro","price":1200}]}`
	bs := []byte(s)

	var people person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("all of the data", people)

	fmt.Println("\nPERSON DATA")
	fmt.Println(people.ID, people.Customer)
	for _, v := range people.Items {
		fmt.Println("\t", v.ID, v.Name, v.Price)
	}
}
