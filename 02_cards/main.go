package main

import "fmt"

func main() {
	cards := []string{"Ace of spades", newCard(), "Six of spades"}

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "King of diamonds"
}
