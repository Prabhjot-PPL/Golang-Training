package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// creating a new type of 'deck'
// which is a slice of string
type deck []string

// function to pass all card values
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// any variable of type deck will now gets access to print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// here we will convert cards into string
// so we can later on implement byte slice in it
// flow -> deck -> slice of string -> string -> byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Function to save data to Hard Drive
func (d deck) saveToFile(filename string) error {
	// three arguments in WriteFile function
	// second argument is byte slice
	// third argument 0666 means anyone can read and write this file
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function to read data from Hard Drive
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Two ways to handle if error occurs
		// Option 1 : log the error and return a call to newDeck() [so we can atleast send a deck to user to work with]
		// Option 2 : Log the error and entirely quit the program
		fmt.Println("Error (reading a file) : ", err)
		os.Exit(1) // value 0 in Exit() indicates success
	}

	// string(bs) = type conversion
	// after type converison, splitting string into slices
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// function to shuffle cards
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		// swaping data in slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
