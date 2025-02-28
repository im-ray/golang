package cards

// go run main.go variables
import (
	"fmt"
)
// Basic Go Types: bool, string, int, float64

// deckSize := 20
// ✅ At the package level, you must use var or const
// ✅ The short variable declaration (:=) is only allowed inside functions.

func VariableDemo() {
	var card string = "Ace of Spades"
	// card_1 := "Ace of Spades" 
	card = "Five of Diamonds" 
	// : is used only once for declaration without explicit mention of types

	card = newCard()
	fmt.Println(card)
}

// Function and return Types
func newCard() string {
	return "Five of Diamonds"
}