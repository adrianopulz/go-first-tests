package main

import "fmt"

func main() {
	deck := newDeck()
	deck.saveToFile("my_deck")
	deck.shuffle()

	firstThreeCards, lastCards := deal(deck, 3)

	fmt.Println("The first three cards of Deck is:")
	firstThreeCards.print()

	fmt.Println("The rest of Deck is:")
	lastCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
