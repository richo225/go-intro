package main

func main() {
	cards := newDeckFromFile()
	cards.shuffle()

	hand := cards.deal(2)
	hand.print()
}
