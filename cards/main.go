package main

func main() {
	cards := newDeckFromFile("./saved_deck.txt")
	cards.shuffle()

	hand := cards.deal(2)
	hand.print()
}
