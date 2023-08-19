package main

import "fmt"

func main() {
	fmt.Println("❌")
}

//build executables for windows, mac, linux
//go build main.go
//GOOS=windows GOARCH=amd64 go build main.go
//GOOS=linux GOARCH=amd64 go build main.go
//GOOS=darwin GOARCH=amd64 go build main.go

