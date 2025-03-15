package main

import (
	"fmt"
	"io"
	"os"
)
// go run main.go myfile.txt will run good but 
// go run main.go main.go--> 
// package command-line-arguments: case-insensitive file name collision: "main.go" and "main.go"


// how to read content of main.go 
// --> go build main.go and then ./main main.go
func main() {
	//fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}