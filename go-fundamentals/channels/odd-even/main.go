package main

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			println("from the even channel:", v)
		case v := <-odd:
			println("from the odd channel:", v)
		case v := <-quit:
			println("from the quit channel:", v)
			return
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- 0
}
