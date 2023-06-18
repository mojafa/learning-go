package main

import "fmt"

type engine struct {
	electric bool
}
type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	// Create two vehicle values using composite literals
	vehicle1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "Model S",
		doors: 4,
		color: "Red",
	}

	vehicle2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 2,
		color: "Blue",
	}

	// Print out each of the vehicle values
	fmt.Println("Vehicle 1:", vehicle1)
	fmt.Println("Vehicle 2:", vehicle2)

	// Print out a single field from each of the vehicle values
	fmt.Println("Vehicle 1 Make:", vehicle1.make)
	fmt.Println("Vehicle 2 Color:", vehicle2.color)

	fmt.Println(vehicle1.electric, vehicle1.make, vehicle1.model)
	fmt.Println(vehicle2.electric, vehicle2.make, vehicle2.model)

	fmt.Println(vehicle1.engine.electric, vehicle1.make, vehicle1.model)
	fmt.Println(vehicle2.engine.electric, vehicle2.make, vehicle2.model)
}

/*
Create a type engine struct, and include this field
electric bool
Create a type vehicle struct, and include these fields
engine
make
model
doors
color
Create two VALUES of TYPE vehicle
use a composite literal
Print out each of these values.
Print out a single field from each of these values.

*/
