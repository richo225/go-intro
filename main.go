package main

func main() {
	cards := deck{"First Card", "Second Card"}
	cards = append(cards, newCard())
	cards.print()
}

func newCard() string {
	return "New card here"
}
