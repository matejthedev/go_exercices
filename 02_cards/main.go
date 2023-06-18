package main

func main() {
	cards := newDeck()

	cards.shuffle()

	// cards.saveToFile("hey")

	// fmt.Println(getDeckFromFile("hey"))

	cards.print()

	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()
}
