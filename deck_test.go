package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	actual := len(newDeck())
	expected := 16

	if actual != expected {
		t.Errorf("expected %q, actual %q", expected, actual)
	}
}

func NewDeckExample() {
	deck := newDeck()
	fmt.Println(len(deck))
	//Output: 16
}

func BenchmarkNewDeck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newDeck()
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("saved_deck_test.txt")
	deck := newDeck()
	deck.saveTofile("saved_deck_test.txt")

	_, error := os.Stat("saved_deck_test.txt")

	actual := os.IsNotExist(error)
	expected := false

	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
