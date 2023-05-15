package main

import "fmt"

func main() {

	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("this is the default case")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	case 44:
		fmt.Println("x is 44")
	case 45:
		fmt.Println("x is 45")
	default:
		fmt.Println("this is the default case")
	}
	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x")
	}
}
