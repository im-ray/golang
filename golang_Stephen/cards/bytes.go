package cards

import "fmt"

// what is byte slice ? [72, 105, 32, 116, 104, 101, 114,]
// ASCII Table will have value corrrsponding to these ASCII Values

func bytes_experiment() {
	greeting := "Hi There!"
	fmt.Println([]byte(greeting))
	// [72 105 32 84 104 101 114 101 33] 
}
