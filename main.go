package main

func main() {
	cards := newDeckFromFile()

	hand := cards.deal(2)
	hand.print()
}
