package main

import (
	"fmt"
	"os"
	"github.com/im-ray/golang/golang_Stephen/cards"

)

// whenever a package main you must also define a function main
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <function_name>")
		return
	}

	choice := os.Args[1] // Get the function name from the command line

	switch choice {
	case "variables":
		cards.VariableDemo()
	default:
		fmt.Println("Unknown function:", choice)
	}
}