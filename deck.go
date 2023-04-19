package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func newDeckFromFile() deck {
	data, err := os.ReadFile("./saved_deck.txt")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(data), ",")
	return deck(s)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) deck {
	return d[:handSize]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveTofile() error {
	return os.WriteFile("./saved_deck.txt", []byte(d.toString()), 0666)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}

	return d
}

// newDeck
// print
// shuffle
// deal
// saveTofile
// newDeckFromFile
