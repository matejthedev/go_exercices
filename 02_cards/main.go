package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("hey")

	fmt.Println(getDeckFromFile("hey"))

	// cards.print()

	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()
}
