package cards

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is slice of strings
type deck []string
// kind of a way to achieve OOP way to extend things 
// think it as an object oriented approach deck class have respective methods
// however since go is not OOP its wrong to say class we call this as receivers
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// by creating a new type with a function that has a receiver we ..
// are adding a method to any value of that type

// d here is a value of type deck

// this or self can be used for receiver code will compile but its not for good practice
// func (this deck) -----


func newDeck() deck {
	// Create and return a list of cards
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four" }

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	hand, remainingCards := d[:handSize], d[handSize:]
	return hand, remainingCards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filname string) deck {
	bs, err := ioutil.ReadFile(filname)
	if err != nil {
		// option 1-> log the error and return a call to newDeck()
		// option 2-> log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	} 
}










