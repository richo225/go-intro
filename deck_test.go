package main

import (
	"fmt"
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
