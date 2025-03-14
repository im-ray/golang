// {"Golang": "Map", "Ruby": "Hash", 
//"JavaScript": "Object", "Python": "Dict" }
package main

// in map {all keys should be of same type, all values should be of same type}
import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	// another way to define map
	var colors_1 map[string]string // zero value of map is empty map
	fmt.Println(colors_1)

	// third way to define map
	colors_2 := make(map[string]string)
	colors_2["white"] = "#ffffff"
	// dont use . to access the map element  for example: colors.white 
	delete(colors_2, "white")
	printMap(colors)

}

//iterating over maps
func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("hex code for color: ", key, ":", val)
	}
}
// how structs and maps are different: there are many but few of them
// map--> reference type, 
// struct--> value type, during compile time we should now all fields



