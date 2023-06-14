package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"spades", "diamonds", "clubs", "hearts"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, hs int) (deck, deck) {
	return d[:hs], d[hs:]
}
