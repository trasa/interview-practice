package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffletest() {

	// not secure
	rand.Seed(time.Now().UnixNano())

	var deck [52]int
	for i := 0; i < len(deck); i++ {
		deck[i] = i
	}
	deck = shuffle(deck)
	for pos, card := range deck {
		fmt.Printf("%d\t%d\n", pos, card)
	}

}

func shuffle(deck [52]int) [52]int {
	// getRandom(floor, ceil) --> random number >= floor and <= ceiling

	for i := len(deck) - 1; i > 1; i-- {
		j := getRandom(0, i)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func getRandom(floor int, ceil int) int {
	return rand.Intn((ceil+1)-floor) + floor
}
