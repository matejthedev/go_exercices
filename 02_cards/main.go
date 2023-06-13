package main

func main() {
	cards := deck{"Ace of spades", newCard(), "Six of spades"}

	cards.print()
}

func newCard() string {
	return "King of diamonds"
}
