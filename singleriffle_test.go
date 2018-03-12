package main

import (
	"fmt"
	"testing"
)

func TestSingleRiffle(t *testing.T) {
	// given shuffledDeck, half1 and half2, determine if half1 and half2 were single-riffled
	// into shuffedDeck.

	full := []int{1, 2, 3, 4, 5}
	one := []int{1, 2, 3, 4}
	two := []int{5}
	if !verifyDeck(full, one, two) {
		panic("expected success")
	}

	full = []int{1, 10, 2, 3, 20, 30, 4, 40}
	one = []int{1, 2, 3, 4}
	two = []int{10, 20, 30, 40}
	if !verifyDeck(full, one, two) {
		panic("expected success")
	}

	badDeck := []int{1, 10, 2, 20, 30, 40, 4, 3}
	one = []int{1, 2, 3, 4}
	two = []int{10, 20, 30, 40}
	if verifyDeck(badDeck, one, two) {
		panic("fail")
	}
}

func verifyDeck(shuffledDeck []int, halfOne []int, halfTwo []int) bool {
	onePos := 0
	twoPos := 0

	for _, card := range shuffledDeck {
		if onePos < len(halfOne) && card == halfOne[onePos] {
			fmt.Printf("Found one: %d (%d)\n", onePos, card)
			onePos++
		} else if twoPos < len(halfTwo) && card == halfTwo[twoPos] {
			fmt.Printf("Found two: %d (%d)\n", twoPos, card)
			twoPos++
		} else {
			// fail!
			fmt.Println("NOT single riffle shuffle!")
			return false
		}
	}
	return onePos == len(halfOne) && twoPos == len(halfTwo)
}

func verifyHalf(shuffledDeck []int, half []int) bool {
	deckPos := 0

	for _, card := range half {
		fmt.Printf("verify %d\n", card)
		success := false
		for deckPos < len(shuffledDeck) {
			fmt.Printf("pos %d: card %d, deck %d\n", deckPos, card, shuffledDeck[deckPos])
			if card == shuffledDeck[deckPos] {
				// found card, success
				success = true
				deckPos++
				break
			}
			deckPos++
		}
		if !success {
			fmt.Println("NOT single riffle")
			return false
		}
	}
	return true
}
