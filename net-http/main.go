package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(writer, "Number of bytes written: %d", n)
	})
	log.Println("Starting server on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
