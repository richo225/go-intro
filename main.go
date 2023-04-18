package main

import (
	"fmt"
)

func main() {
	cards := []string{"First Card", "Second Card"}
	cards = append(cards, newCard())

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "New card here"
}
