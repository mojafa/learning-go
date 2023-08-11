package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(dowork)

}

func dowork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
