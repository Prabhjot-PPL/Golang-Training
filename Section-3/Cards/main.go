package main

func main() {

	// using newDeck function(deck.go) to pass all card values

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// READING A FILE FROM HARD DRIVE
	// cards := newDeck()
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// SAVING TO HARD DRIVE
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// CHECKING DEAL FUNCTION, SPLIT DECK IN TWO PARTS
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("------------------")
	// remainingCards.print()

	// BYTE CONVERSION DEMO
	// greeting := "hi there"
	// fmt.Println([]byte(greeting))
}
