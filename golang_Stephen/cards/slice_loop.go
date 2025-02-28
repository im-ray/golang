package cards

import "fmt"

// Array --> Fixed length
// Slice --> an array that can grow or shrink
// array and slice both must be defined with a data types : Homogenous

func cardslice() {
	cards := []string{"Ace of Diamonds", newCard_1()}
	
	// append function returns a new slice
	cards = append(cards, "Six of Spades")

	// iterate over slice
	for index , card := range cards {
		fmt.Println(index, card)
	}
}


// Every Variable we declare must be used in our code

func newCard_1() string {
	return "Five of Diamonds"
}
// slices are zero indexed
// slicing in slices similiar to as in python


