//logical oprators

package main

func main() {
	// && and
	// || or
	// ! not

	x := 45
	y := 42

	if x < 42 && y < 42 {
		println("001")
	}
	if x > 30 || y < 42 {
		println("002")
	}

	if !(x < 42) {
		println("003")
	}

	if x != y {
		println("004")
	}
	if x == y {
		println("005")
	}

	if !(x < 42 && y < 42) {
		println("006")
	}

	if !(x > 42 || y < 42) {
		println("007")
	}
}
