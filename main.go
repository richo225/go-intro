package main

func main() {
	cards := newDeck()

	hand := cards.deal(4)
	hand.print()
	hand.saveTofile()
}
